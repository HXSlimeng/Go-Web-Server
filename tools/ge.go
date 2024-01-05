package tools

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func GeTemplate(pkgName string) {
	wg := new(sync.WaitGroup)

	 filePath := fmt.Sprintf("app/%s",pkgName)
	if _,err := os.Stat(filePath);os.IsNotExist(err) {
		
		//创建文件夹
		err := os.Mkdir(filePath,os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}

		//创建文件
		fileList := []string{"service","controller","model"}
		for _, v := range fileList {
			wg.Add(1)
			go func(v string){
				toCreatePath := fmt.Sprintf("%s/%s.%s.go",filePath,pkgName,v)
			
				file, err := os.Create(toCreatePath)
				if  err == nil{
					file.WriteString(fmt.Sprintf("package %s \n",pkgName))
				}
				
				defer file.Close()
				defer wg.Done()
			}(v)
		} 
	}else{
		log.Fatal(filePath + " 文件夹已存在!")
	}
	wg.Wait()
}