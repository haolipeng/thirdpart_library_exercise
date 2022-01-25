package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	appv1 "k8s.io/api/apps/v1"
	"k8s.io/api/apps/v1beta2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//枚举集群中pod信息
	ListPods(clientset)

	//使用ds.yaml文件来部署DaemonSet pod
	DaemonSetWithYaml(clientset, "add", "kube-system")
	DaemonSetWithYaml(clientset, "get", "kube-system")
	DaemonSetWithYaml(clientset, "delete", "kube-system")
}

/*
DaemonSetWithYaml
clientset: 和k8s集群的客户端连接
action:操作动作
	add:添加
	update:更新
	delete:删除
*/
func DaemonSetWithYaml(clientset *kubernetes.Clientset, action string, namespace string) error {
	//解析yaml文件
	dir, _ := os.Getwd()
	fmt.Println(dir)
	fileName := dir + "\\" + "ds.yaml"

	yamlContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("ioutil.ReadFile failed\n")
	}

	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode(yamlContent, nil, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error while decoding YAML object. Err was:%s", err))
	}

	var ds *appv1.DaemonSet
	switch o := obj.(type) {
	case *v1.Pod:
		fmt.Println("pod type")
	case *v1beta1.DaemonSet:
		fmt.Println("v1 beta1.DaemonSet type:", o.Name)
	case *v1beta2.DaemonSet:
		fmt.Println("v1 beta2.DaemonSet type:", o.Name)
	case *appv1.DaemonSet:
		ds = o
		fmt.Println("appv1.DaemonSet type:", o.Name)
	default:
		fmt.Println("unsupported type")
		return errors.New("unsupported type yaml file")
	}

	//测试下连续添加两次会如何?
	if ds != nil {
		//部署daemonset
		switch action {
		case "add":
			dsNew, err := clientset.AppsV1().DaemonSets(namespace).Create(context.TODO(), ds, metav1.CreateOptions{})
			if err != nil {
				fmt.Printf("name:%s\n status:%s", dsNew.Name, dsNew.Status)
				return err
			}
		case "delete":
			//传递要删除的DaemonSet名称
			err := clientset.AppsV1().DaemonSets(namespace).Delete(context.TODO(), ds.Name, metav1.DeleteOptions{})
			if err != nil {
				fmt.Printf("delete DaemonSets failed:%s\n", ds.Name)
				return err
			}
		case "get":
			dsGet, err := clientset.AppsV1().DaemonSets(namespace).Get(context.TODO(), ds.Name, metav1.GetOptions{})
			if err != nil {
				fmt.Printf("get DaemonSets failed:%s\n", dsGet.Name)
				return err
			}
		case "update":
			dsUpdate, err := clientset.AppsV1().DaemonSets(namespace).Update(context.TODO(), ds, metav1.UpdateOptions{})
			if err != nil {
				fmt.Printf("update DaemonSets failed:%s\n", dsUpdate.Name)
				return err
			}
		}

	}
	return err
}

func ListPods(clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal("pods list failed")
	}

	for _, pod := range pods.Items {
		//fmt.Printf("Name: %s, Status: %s Namespace:%s NodeName:%s\n",
		//	pod.ObjectMeta.Name, pod.Status.Phase, pod.Namespace, pod.Spec.NodeName)
		fmt.Printf("pod %s in Node name:%s\n", pod.Name, pod.Spec.NodeName)
		/*if pod.ObjectMeta.Name == "my-centos" {
			containers := pod.Spec.Containers
			fmt.Printf("pod UUID:%s NodeName:%s\n", pod.ObjectMeta.UID, pod.Spec.NodeName)
			for _, c := range containers {
				fmt.Printf("container name:%s", c.Name)
			}
		}*/
	}
}

func WatchPods() {
	//Watch监视功能的使用
	//service -> lables
	/*labels := "system=centos"
	result, err := clientset.CoreV1().Pods("default").Watch(context.TODO(), metav1.ListOptions{LabelSelector: labels})
	if err != nil {
		panic(err.Error())
	}
	resultCh := result.ResultChan()
	for ch := range resultCh {
		fmt.Printf("type:%s\n", string(ch.Type))
	}*/
}

func ExecCmdExample(client kubernetes.Interface, config *restclient.Config, podName string,
	command string /*, stdin io.Reader, stdout io.Writer, stderr io.Writer*/) error {
	/*cmd := []string{
		"sh",
		"-c",
		command,
	}*/
	var stdin bytes.Buffer
	req := client.CoreV1().RESTClient().Post().Resource("pods").Name(podName).
		Namespace("default").SubResource("exec")
	option := &v1.PodExecOptions{
		Command: strings.Fields(command),
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:     true,
	}
	/*if stdin == nil {
		option.Stdin = false
	}*/
	req.VersionedParams(
		option,
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return err
	}

	//stdin,stdout,stderr
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  &stdin,
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		return err
	}
	fmt.Printf("output:%s err:%s\n", stdout.String(), stderr.String())

	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
