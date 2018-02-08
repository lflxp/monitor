// +build linux

package deploy

import (
	"fmt"
	"io"
	"os"

	"github.com/lflxp/monitor/utils"
)

func DeployService() error {
	var err error
	path := "/etc/systemd/system/monitor.service"
	services := `[Unit]
	Description=monitor
	After=network.target
	
	[Service]
	Type=oneshot
	ExecStart=/etc/init.d/monitor start
	ExecStop=/etc/init.d/monitor stop
	ExecReload=/etc/init.d/monitor restart
	RemainAfterExit=yes
	
	[Install]
	WantedBy=multi-user.target
	`
	if utils.CheckFileLsExist(path) {
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	utils.Check(err)
	n, err := io.WriteString(f, services) //写入文件(字符串)
	utils.Check(err)
	fmt.Printf("写入 %d 个字节n", n)
	return err
}

func DeployInited() error {
	var err error
	path := "/etc/init.d/monitor"
	services := `#!/bin/sh
	#
	export LANG=en_US.UTF-8
	
	SERVICE_NAME="nta"
	APP_PATH="/opt/app"
	EXEC=$APP_PATH/$SERVICE_NAME/$SERVICE_NAME
	PIDS="/opt/data/pids"
	PID_FILE=$PIDS/$SERVICE_NAME
	LOG_PATH="/opt/data/logs"
	ERROR_FILE="error"
	LOG_FILE="log"
	
	case $1 in
		start)
			echo "Service $SERVICE_NAME Starting ..."
			if [ ! -f $PID_FILE ]; then
				cd $APP_PATH/$SERVICE_NAME 
				nohup $EXEC 2>> $LOG_PATH/$SERVICE_NAME/$ERROR_FILE >> $LOG_PATH/$SERVICE_NAME/$LOG_FILE &
				echo $! > $PID_FILE
				echo "Service $SERVICE_NAME started PID is `cat $PID_FILE`"
			else
				echo "Service $SERVICE_NAME is already running PID is `cat $PID_FILE`"
			fi
		;;
		stop)
			if [ -f $PID_FILE ]; then
				PID=$(cat $PID_FILE);
				echo "Service $SERVICE_NAME stoping PID is `cat $PID_FILE`"
				kill $PID;
				echo "Service $SERVICE_NAME stopped "
				rm $PID_FILE
			else
				echo "Service $SERVICE_NAME is not running ..."
			fi
		;;
		restart)
			if [ -f $PID_FILE ]; then
				PID=$(cat $PID_FILE);
				echo "Service $SERVICE_NAME stopping PID is `cat $PID_FILE`";
				kill $PID;
				echo "Service $SERVICE_NAME stopped ";
				rm -f $PID_FILE
				echo "Service $SERVICE_NAME starting ..."
				nohup $EXEC 2>> $LOG_PATH/$SERVICE_NAME/$ERROR_FILE > $LOG_PATH/$SERVICE_NAME/$LOG_FILE &
				echo $! > $PID_FILE
				echo "Service $SERVICE_NAME started PID is `cat $PID_FILE`"
			else
				echo "Service $SERVICE_NAME is not running ..."
			fi
		;;
	esac`
	if utils.CheckFileLsExist(path) {
		f, err = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	utils.Check(err)
	n, err := io.WriteString(f, services) //写入文件(字符串)
	utils.Check(err)
	fmt.Printf("写入 %d 个字节n", n)
	return err	
}

func Start() {
	reload := "systemctl daemon-reload"
	status := "systemctl status monitor"
	start := "systemctl start monitor"
	log.Println(utils.ExecCommand(reload))
	log.Println(utils.ExecCommand(status))
	log.Println(utils.ExecCommand(start))
}