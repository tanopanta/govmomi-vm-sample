package main

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/examples"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func main() {
	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		fmt.Println("from name")

		f := find.NewFinder(c)
		vm, err := f.VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		fmt.Printf("name: %s, path: %s, mo-id: %v\n", vm.Name(), vm.InventoryPath, vm.Reference())

		return nil
	})

	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		fmt.Println("from folder name")

		f := find.NewFinder(c)
		vms, err := f.VirtualMachineList(ctx, "/DC0/vm/*")
		if err != nil {
			return err
		}

		for _, vm := range vms {
			fmt.Printf("name: %s, path: %s, mo-id: %v\n", vm.Name(), vm.InventoryPath, vm.Reference())
		}

		return nil
	})

	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		fmt.Println("from mo-id")

		f := find.NewFinder(c)
		ref, err := f.ObjectReference(ctx, types.ManagedObjectReference{
			Type:  "VirtualMachine",
			Value: "vm-54",
		})
		if err != nil {
			return err
		}
		vm, _ := ref.(*object.VirtualMachine)

		fmt.Printf("name: %s, path: %s, mo-id: %v\n", vm.Name(), vm.InventoryPath, vm.Reference())

		return nil
	})

	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		fmt.Println("from mo-id(without finder)")

		vm := object.NewVirtualMachine(c, types.ManagedObjectReference{
			Type:  "VirtualMachine",
			Value: "vm-54",
		})

		name, err := vm.ObjectName(ctx)
		if err != nil {
			return err
		}
		path, err := find.InventoryPath(ctx, c, vm.Reference())
		if err != nil {
			return err
		}

		fmt.Printf("name: %s, path: %s, mo-id: %v\n", name, path, vm.Reference())

		return nil
	})

	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		fmt.Println("from vm uuid")

		si := object.NewSearchIndex(c)
		ref, err := si.FindByUuid(
			ctx,
			nil, // use default dc
			"265104de-1472-547c-b873-6dc7883fb6cb",
			true,
			nil, // if Instance UUID => types.NewBool(true)
		)
		if err != nil {
			return err
		}

		vm, _ := ref.(*object.VirtualMachine)

		name, err := vm.ObjectName(ctx)
		if err != nil {
			return err
		}
		path, err := find.InventoryPath(ctx, c, vm.Reference())
		if err != nil {
			return err
		}

		fmt.Printf("name: %s, path: %s, mo-id: %v\n", name, path, vm.Reference())

		return nil
	})

	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		fmt.Println("from host")

		f := find.NewFinder(c)
		host, err := f.HostSystem(ctx, "DC0_H0")
		if err != nil {
			return err
		}

		var m mo.HostSystem
		err = host.Properties(ctx, host.Reference(), []string{"vm"}, &m)
		if err != nil {
			return err
		}

		vmrefs := m.Vm
		for _, vmref := range vmrefs {
			vm := object.NewVirtualMachine(c, vmref)

			name, err := vm.ObjectName(ctx)
			if err != nil {
				return err
			}
			path, err := find.InventoryPath(ctx, c, vm.Reference())
			if err != nil {
				return err
			}

			fmt.Printf("name: %s, path: %s, mo-id: %v\n", name, path, vm.Reference())
		}

		return nil
	})
}
