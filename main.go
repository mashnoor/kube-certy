package main

import (
	"fmt"
	"os/exec"

	"strings"
	"errors"
)

func printOut(cmd *exec.Cmd) error  {
	stdout, err := cmd.Output()
	if err == nil {
		fmt.Println(string(stdout))
		if strings.Contains(string(stdout), "Error") {
			return errors.New("empty name")
		}
	} else {
		fmt.Println(err)
		return nil
	}
	return nil

}

func retry() {
	fmt.Println("Retrying...")

	totalError := 9
	for totalError > 0 {
		totalError = 0
		cmd := exec.Command("kubeadm", "certs", "renew", "apiserver-kubelet-client")
		err := printOut(cmd)
		if err != nil {
			totalError ++
		}


		cmd = exec.Command("kubeadm", "certs", "renew", "apiserver")
		err = printOut(cmd)
		if err != nil {
			totalError ++
		}

		cmd = exec.Command("kubeadm", "certs", "renew", "admin.conf")
		err = printOut(cmd)
		if err != nil {
			totalError ++
		}

		fmt.Println(totalError)
		// time.Sleep(5 * time.Second)

	}




}

func main()  {
	cmd := exec.Command("mv", "/etc/kubernetes/pki/ca.crt", "/etc/kubernetes/pki/ca.crt.bk")
	printOut(cmd)
	cmd = exec.Command("mv", "/etc/kubernetes/pki/ca.key", "/etc/kubernetes/pki/ca.key.bk")
	printOut(cmd)

	cmd = exec.Command("kubeadm", "init", "phase", "certs", "ca")
	_ = printOut(cmd)

	retry()





}
