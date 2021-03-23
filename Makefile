

gen.mono.repo:
	cd tmp/; cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="mono-repo/go-zero"

gen.basic.app:
	cd tmp/mall/app/basic; cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app/go-zero"

gen.biz.app:
	cd tmp/mall/app/biz; cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app/go-zero"
