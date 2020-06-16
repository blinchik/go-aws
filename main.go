package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"

	mEC2 "github.com/blinchik/go-aws/lib/manage-ec2"
	hostRe "github.com/blinchik/go-aws/lib/refresh-hostnames"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	raw := flag.Bool("raw", false, "raw")
	describe := flag.Bool("d", false, "describe")
	freeTag := flag.Bool("f", false, "free tag")
	vpc := flag.Bool("vpc", false, "vpc")
	sg := flag.Bool("sg", false, "sg")
	sb := flag.Bool("sb", false, "subnet")
	vol := flag.Bool("vol", false, "volume")
	importkey := flag.Bool("importkey", false, "importkey")

	flag.Parse()

	if *raw {

		if os.Args[2] == "start" {

			summary := mEC2.DescribeByOperationTag(os.Args[3])
			mEC2.StartEC2(summary.InstanceId)
			return
		}
	} else {
		if os.Args[1] == "start" {
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

	if *describe {
		if *freeTag {

			summary := mEC2.DescribeByGeneralTag(os.Args[3], os.Args[4])

			fmt.Println(summary)

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

	if *importkey {

		name := os.Args[2]

		mEC2.ImportKey(name)

	}

	if *describe {
		if *vpc {

			mEC2.VpcDescribe()

		}
	}

	if *describe {
		if *sg {

			mEC2.SgDescribe()

		}
	}

	if *describe {
		if *sb {

			mEC2.SubnetDescribe()

		}
	}

	if *vol {
		if *sg {

			mEC2.DescribeAllVols()

		}
	}

	if os.Args[1] == "refresh" {
		hostRe.HostnamesRefresh()

	}

	if os.Args[1] == "ebsList" {
		mEC2.DescribeAllVols()

	}

}
