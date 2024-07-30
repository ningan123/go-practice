package main

import (
	"context"
	"fmt"
	"log"
	"syscall"
	"time"

	// "time"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
)

func main() {
	if err := containerExample(); err != nil {
		log.Fatal(err)
	}
}

func containerExample() error {
	// 初始化client
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()

	// 拉取镜像
	ctx := namespaces.WithNamespace(context.Background(), "default")
	// image, err := client.Pull(ctx, "docker.io/library/golang:1.20", containerd.WithPullUnpack)
	// image, err := client.Pull(ctx, "docker.io/library/busybox:1.36", containerd.WithPullUnpack)
	image, err := client.Pull(ctx, "docker.io/library/redis:alpine", containerd.WithPullUnpack)
	if err != nil {
		return err
	}
	log.Printf("Successfully pulled %s image\n", image.Name())

	// 创建OCI Spec
	// container, err := client.NewContainer(
	// 	ctx,
	// 	"busybox",
	// 	containerd.WithNewSnapshot("busybox", image),
	// 	containerd.WithNewSpec(
	// 		oci.WithImageConfig(image),
	// 		// oci.WithProcessArgs("sleep", "infinity"),
	// 	),
	// )
	container, err := client.NewContainer(
		ctx,
		"redis-server",
		containerd.WithNewSnapshot("redis-server-snapshot", image),
		containerd.WithNewSpec(
			oci.WithImageConfig(image),
			// oci.WithProcessArgs("sleep", "infinity"),
		),
	)
	if err != nil {
		return err
	}
	defer container.Delete(ctx, containerd.WithSnapshotCleanup)
	log.Printf("Successfully created container with ID %s and snapshot with ID redis-server", container.ID())
	// time.Sleep(60 * time.Second)

	// 创建task
	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	if err != nil {
		return err
	}
	log.Printf("Successfully created task %s", task.ID())
	defer task.Delete(ctx)

	// 启动task
	exitStatusC, err := task.Wait(ctx)
	if err != nil {
		return err
	}
	err = task.Start(ctx)
	if err != nil {
		return err
	}
	log.Printf("Successfully started task %s", task.ID())

	// 停止task
	log.Printf("\n\n##====== redis server启动后输出的日志 ======##")
	time.Sleep(10 * time.Second)
	log.Printf("\n\n##====== redis进程处理sigterm信号 ======##")
	err = task.Kill(ctx, syscall.SIGTERM)
	if err != nil {
		return err
	}

	status := <-exitStatusC
	code, exitedAt, err := status.Result()
	if err != nil {
		return err
	}
	log.Printf("\n\n##====== 获取到redis进程的退出状态码 ======##")
	fmt.Printf("busybox exited with status: %d, exitedAt: %v\n", code, exitedAt)
	return nil
}
