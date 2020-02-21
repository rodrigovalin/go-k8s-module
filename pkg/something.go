package pkg

import (appsv1 "k8s.io/api/apps/v1")

type ExportedStruct struct {
	Name string
	Sts appsv1.StatefulSet
}
