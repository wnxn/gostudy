package main

import (
	"context"
	"flag"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"k8s.io/klog"
	"os"
)

var (
	endpoint = flag.String("endpoint", "unix:///tmp/csi.sock", "CSI endpoint")
)

func main() {
	conn, err := grpc.Dial(*endpoint,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{PermitWithoutStream: true}))
	if err != nil{
		klog.Errorf("error connecting to CSI driver: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	DeleteSnapshot(csi.NewControllerClient(conn))

}

func DeleteSnapshot(cc csi.ControllerClient){
	volReq:=&csi.CreateVolumeRequest{
		Name:                      "csi-client",
		VolumeCapabilities: []*csi.VolumeCapability{
			&csi.VolumeCapability{
				AccessType:           nil,
				AccessMode:          &csi.VolumeCapability_AccessMode{
					Mode:                 csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
				},
			},
		},
	}
	volResp, err := cc.CreateVolume(context.Background(),volReq)
	if err != nil{
		klog.Error(err)
	}
	snapReq := &csi.CreateSnapshotRequest{
		SourceVolumeId:       volResp.Volume.VolumeId,
		Name:                 "csi-client",
	}
	snapResp, err := cc.CreateSnapshot(context.Background(), snapReq)
	if err != nil{
		klog.Error(err)
	}
	snapDelReq := &csi.DeleteSnapshotRequest{
		SnapshotId:           snapResp.Snapshot.SnapshotId,
	}
	snapDelResp, err := cc.DeleteSnapshot(context.Background(), snapDelReq)
	if err != nil{
		klog.Error(err)
	}
	klog.Info(snapDelResp)
	volDelReq := &csi.DeleteVolumeRequest{
		VolumeId:             volResp.Volume.GetVolumeId(),
	}
	volDelResp, err := cc.DeleteVolume(context.Background(), volDelReq)
	if err != nil{
		klog.Error(err)
	}
	klog.Info(volDelResp)
}

