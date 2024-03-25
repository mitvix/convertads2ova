#!/bin/bash
# Script para conversão de versão do vAppliance AWS Application Discovery Service
# para sistemas VMware legados como VMware ESXi 5.5 e inferiores
#
# Antes de continuar baixe a última versão do VMware-ovftool-XXX-lin.x86_64.bundle.
# https://customerconnect.vmware.com/downloads/get-download?downloadGroup=OVFTOOL441
# 2024/03/23 - convertadsova.sh v0.2 <alexmanfrin@gmail.com>
#
# Use ./convertadsova.sh nome_do_virtual_appliance.ova

# checagem básica de erros
if [ ! -f "$1" ]; then
    echo "Arquivo $1 não encontrado."
    echo "Use: $0 nome_do_virtual_appliance.ova"
    exit 1
# termina se ovftool não for encontrado
elif ! command -v ovftool &> /dev/null 
then
    echo "<ovftool> não foi encontrado"
    echo "Faça o download do ovftool do site da VMware antes de continuar!"
    exit 1
fi

# init
fr_version=11
to_version=10
filename=${1%.*}
VMX=$filename".vmx"
OVA=$filename".ova"
VMDK=$filename"-disk1.vmdk"
NEWOVA=$filename"_new.ova"

# valida a existência do filename e executa conversão
if [ -f "$1" ]; then
    echo "Convertendo vAppliance, aguarde..."
    ovftool $OVA $VMX
    
    echo "Convertendo versão do hardware..."
    sed -i -e 's/virtualhw.version = "'$fr_version'"/virtualhw.version = "'$to_version'"/' $VMX
    
    echo "Convertendo appliance com hash SHA1..."
    ovftool --shaAlgorithm=SHA1 $VMX $NEWOVA
    
    echo "Realizando limpeza..."
    if [ -f $VMDK ] || [ -f $VMX ]; then
        rm -fv $VMX
        rm -fv $VMDK
    fi

    # checagem de erro durante conversão
    case $? in
        0) echo "Processo concluído com sucesso, siga com a importação do $NEWOVA no vSphere!"
            ls -has $NEWOVA            
        ;;
        *) echo "Erro: Foi encontrado uma falha durante a conversão! [$?]"
            exit 1
        ;;
    esac
fi
