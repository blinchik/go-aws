package main

import (
	"fmt"
	"os"
	"flag"
	"reflect"

	mEC2 "github.com/blinchik-io/go-aws/lib/manage-ec2"
	hostRe "github.com/blinchik-io/go-aws/lib/refresh-hostnames"
)

func main() {

	if os.Args[1] == "start" {
		
		raw := flag.Bool("raw", false, "raw")
		flag.Parse()
		
		if *raw{
			summary := mEC2.DescribeByOperationTag(os.Args[2])
			mEC2.StartEC2(summary.InstanceId)
			return
		}else{
			
		summary := mEC2.DescribeByOperationTag(os.Args[2])

		mEC2.StartEC2(summary.InstanceId)

		hostRe.HostnamesRefresh()
			
		}




	}

	if os.Args[1] == "stop" {

		summary := mEC2.DescribeByOperationTag(os.Args[2])

		mEC2.StopEC2(summary.InstanceId)
	}

	if os.Args[1] == "-d" {
		if os.Args[2] == "-f" {

			summary := mEC2.DescribeByGeneralTag("Function", os.Args[3])

			v := reflect.ValueOf(&summary).Elem()

			typeOfS := v.Type()

			for z := 0; z < len(summary.InstanceId); z++ {

				fmt.Printf("\n")
				fmt.Printf("========================================================================")
				fmt.Printf("\n")

				for i := 0; i < v.NumField(); i++ {

					vv := v.Field(i).Interface()
					if len(vv.([]*string))-z != 0 {

						fmt.Printf("%-25s %-25s\n", typeOfS.Field(i).Name, *vv.([]*string)[z])

					}
				}

				fmt.Printf("========================================================================")
				fmt.Printf("\n")

			}

		}

	}

	if os.Args[1] == "-d" {
		if os.Args[2] == "-t" {

			summary := mEC2.DescribeByGeneralTag(os.Args[3], os.Args[4])
			v := reflect.ValueOf(&summary).Elem()

			typeOfS := v.Type()

			for z := 0; z < len(summary.InstanceId); z++ {

				fmt.Printf("\n")
				fmt.Printf("========================================================================")
				fmt.Printf("\n")

				for i := 0; i < v.NumField(); i++ {

					vv := v.Field(i).Interface()
					if len(vv.([]*string))-z != 0 {

						fmt.Printf("%-25s %-25s\n", typeOfS.Field(i).Name)

					}
				}

				fmt.Printf("========================================================================")
				fmt.Printf("\n")

			}

		}

	}

	if os.Args[1] == "refresh" {
		hostRe.HostnamesRefresh()

	}

	if os.Args[1] == "ebsList" {
		mEC2.DescribeAllVols()

	}

}
