/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ec2igroupv1 "ec2i/api/v1"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/ec2"
)

// EC2IReconciler reconciles a EC2I object
type EC2IReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=ec2igroup.ec2idomain,resources=ec2is,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ec2igroup.ec2idomain,resources=ec2is/status,verbs=get;update;patch

func (r *EC2IReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("ec2i", req.NamespacedName)

	// your logic here

	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-east-2")},
	// )

	//svc := ec2.New(sess)

	var ec2I ec2igroupv1.EC2I

	if err := r.Get(ctx, req.NamespacedName, &ec2I); err != nil {
		//log.Error(err, "unable to fetch EC2Instance")
		fmt.Println("unable to fetch EC2Instance")
	} else {
		// fmt.Println(ec2I.Spec.ImageId)
		// fmt.Println(ec2I.Spec.InstanceType)
		fmt.Println("Doing nothing")
		fmt.Println(ec2I.Spec.Affinity)
	}

	// runR, err := svc.RunInstances(&ec2.RunInstancesInput{

	// 	ImageId:      aws.String(ec2I.Spec.ImageId),
	// 	InstanceType: aws.String(ec2I.Spec.InstanceType),
	// 	MinCount:     aws.Int64(1),
	// 	MaxCount:     aws.Int64(1),
	// })

	// if err != nil {
	// 	fmt.Println("\nCould not create instance")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Created instance")
	// 	fmt.Println(runR)
	// }

	return ctrl.Result{}, nil
}

func (r *EC2IReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ec2igroupv1.EC2I{}).
		Complete(r)
}
