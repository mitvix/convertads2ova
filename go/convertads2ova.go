/* golang go
 * Script para conversão de versão do vAppliance ADS para sistemas VMware legados
 * Antes de continuar baixe a última versão do VMware-ovftool(Windows||Linux)
 * Use go run convertads2ova.go Nome_do_vAppliance.ova
 *
 */

package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func checkFileExists(filepath string) bool {
	_, error := os.Stat(filepath)
	return !errors.Is(error, os.ErrNotExist)
}

func executeCommand(command string, args ...string) bool {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return false
	}
	fmt.Println(out.String())
	return true
}

func main() {
	var execute bool
	args := os.Args

	if len(args[1:]) < 1 {
		fmt.Printf("Use: %v nome_do_virtual_appliance.ova\n", args[0:])
		return
	}
	filename := strings.Join(args[1:2], "")
	isFileExists := checkFileExists(filename)
	if !isFileExists {
		fmt.Printf("Arquivo %v não encontrado\n", filename)
		return
	}

	fr_version := "virtualhw.version = \"11\""
	to_version := "virtualhw.version = \"10\""
	filename = strings.Split(filename, ".")[0]
	vmx := filename + ".vmx"
	ova := filename + ".ova"
	vmdk := filename + "-disk1.vmdk"
	newova := filename + "_new.ova"
	newhash := "--shaAlgorithm=SHA1"

	ovftool := exec.Command("ovftool")
	err := ovftool.Start()
	if err != nil {
		fmt.Println("<ovftool> não foi encontrado\nFaça o download do ovftool do site da VMware antes de continuar!")
		fmt.Println("https://customerconnect.vmware.com/downloads/get-download?downloadGroup=OVFTOOL441")
		return
	}

	fmt.Println("Convertendo vAppliance, aguarde...")
	execute = executeCommand("ovftool", ova, vmx)
	if !execute {
		log.Fatalf("Erro durante a conversão %s para %s\n", ova, vmx)
	}

	fmt.Println("Convertendo a versão do hardware...")
	vmxExist := checkFileExists(vmx)
	if !vmxExist {
		log.Fatalf("Erro durante conversão do hardware. Arquivo %s não existe", vmx)
	}

	input, err := os.ReadFile(vmx)
	if err != nil {
		log.Fatalln(err)
	}

	output := bytes.Replace(input, []byte(fr_version), []byte(to_version), -1)

	if err = os.WriteFile(vmx, []byte(output), 0644); err != nil {
		log.Fatalln(err)
	}
	//fmt.Printf("%v", string(output[:]))

	fmt.Println("Convertendo appliance com hash SHA1...")
	vmdkExist := checkFileExists(vmdk)
	if !vmdkExist {
		log.Fatalf("Erro durante conversão. Arquivo %s não existe", vmdk)
	}

	execute = executeCommand("ovftool", newhash, vmx, newova)
	if !execute {
		log.Fatalf("Erro durante a conversão com hash SHA1!")
	}

	fmt.Println("Realizando limpeza...")
	r_vmx := os.Remove(vmx)
	r_vmdk := os.Remove(vmdk)
	if r_vmx != nil || r_vmdk != nil {
		log.Fatalf("%v\n %v\n", r_vmx, r_vmdk)
	}
	fmt.Println("Processo concluído com sucesso, siga com a importação do ", newova, "no VMware vSphere.")
}
