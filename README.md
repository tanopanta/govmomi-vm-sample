# govmomi-vm-sample
vm get sample(finder, serach index etc..)

## output
```sh
> go run main.go
from name
name: DC0_H0_VM0, path: /DC0/vm/DC0_H0_VM0, mo-id: VirtualMachine:vm-54
from folder name
name: DC0_H0_VM0, path: /DC0/vm/DC0_H0_VM0, mo-id: VirtualMachine:vm-54
name: DC0_H0_VM1, path: /DC0/vm/DC0_H0_VM1, mo-id: VirtualMachine:vm-57
name: DC0_C0_RP0_VM0, path: /DC0/vm/DC0_C0_RP0_VM0, mo-id: VirtualMachine:vm-60
name: DC0_C0_RP0_VM1, path: /DC0/vm/DC0_C0_RP0_VM1, mo-id: VirtualMachine:vm-63
from mo-id
name: DC0_H0_VM0, path: /DC0/vm/DC0_H0_VM0, mo-id: VirtualMachine:vm-54
from mo-id(without finder)
name: DC0_H0_VM0, path: /DC0/vm/DC0_H0_VM0, mo-id: VirtualMachine:vm-54
from vm uuid
name: DC0_H0_VM0, path: /DC0/vm/DC0_H0_VM0, mo-id: VirtualMachine:vm-54
from host
name: DC0_H0_VM0, path: /DC0/vm/DC0_H0_VM0, mo-id: VirtualMachine:vm-54
name: DC0_H0_VM1, path: /DC0/vm/DC0_H0_VM1, mo-id: VirtualMachine:vm-57
```