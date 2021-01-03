# demo-ext-apiserver

## prerequisites
apiserver-boot version = v1.18.0

docs: https://github.com/kubernetes-sigs/apiserver-builder-alpha/blob/v1.18.0/docs/tools_user_guide.md


## init
```sh
DIR=$GOPATH/src/demo.com/ext-apiserver
mkdir -p $DIR
cd $DIR

cat <<EOF >boilerplate.go.txt
/*
Copyright YEAR The Kubernetes Authors.

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
EOF

apiserver-boot  init repo --domain demo.com/ext-apiserver

apiserver-boot create group version resource --group pet --version v1 --kind Cat
```
