![Test Status](https://github.com/code-yeongyu/gin-psql-microservice-template/actions/workflows/test.yml/badge.svg)

# gin-psql-microservice-template

## 프로젝트 구조

-   cmd/server
    -   메인 프로젝트 디렉토리
    -   그 안의 구조 생략
-   configs
    -   환경변수 같은 설정 파일
-   test
    -   테스트 파일

## 모듈 이름 변경

이 프로젝트는 go mod를 사용합니다.
모듈 명(패키지 명)은 gin_psql_microservice_template 로 되어있고, 이를 변경하기 위한 파이썬 스크립트를 제공합니다.  
macOS 환경에서 테스트 되었습니다.

```sh
python3 rename.py <new_name>
```

## 환경 변수

실행시 필요한 환경 변수는 [configs/envs.go](./configs/envs.go) 에서 확인 할 수 있습니다.  
또한 godotenv 모듈을 사용하기 때문에, .env 파일을 활용 할 수 있습니다.

## <a name="execution"></a>실행

의존성 패키지들을 다운 받기 위해 다음의 명령을 실행해주세요.

```sh
go mod tidy
```

다음의 명령어로 프로젝트를 실행 할 수 있습니다.

```sh
make run
```

## 빌드

### 리눅스 바이너리 빌드

```sh
make build
```

### 그 외 운영체제 바이너리 빌드

```sh
go build -ldflags="-s -w" -o bin/main cmd/server/main.go
```

### 도커 빌드

빌드 프로세스에는 멀티스테이지 빌드가 사용됩니다.  
golang:1.16.2-alpine3.13 이미지를 통해 바이너리 빌드를 진행합니다.  
빌드된 이미지는 빈 alpine

```sh
docker build -t <image_name> .
```

## API 문서

### 실행하여 문서를 얻기

[실행](#execution) 후 /swagger/index.html 로 이동하면, api들을 테스트 할 수 있습니다.  
실행 포트를 변경하지 않았다면 <http://localhost:8080/swagger/index.html>에서 접근 할 수 있습니다.

### 직접 문서 파일 얻기

실행하지 않고 직접 json이나 yaml로 된 문서 파일을 얻고 싶다면, [cmd/server/docs](./cmd/server/docs) 에서 docs.json과 docs.yaml을 얻을 수 있습니다.

### 문서 업데이트 하기

본 프로젝트의 문서화는 swagger, gin-swagger 을 통해 자동화된 방식으로 진행되었습니다.  
코드가 변경되었다면 다음의 명령어로 변경 사항을 저장 할 수 있습니다.

```sh
make docs
```
