# Decort SDK

Decort SDK - это библиотека, написанная на языке GO, позволяющая взаимодействовать с API облачной платформы **DECORT**. Библиотека содеражит в себе структуры и методы, необходимые для отправки запросов. Decort SDK имеет встроенный http-клиент и поддерживает разные способы авторизации на платформе. Библиотека так же содержит в себе модели ответов от платформы.

## Версии

    - Версия 1.0.x Decort-SDK соответствует 3.8.4 версии платформы
    - Версия 1.1.x Decort-SDK соответствует 3.8.5 версии платформы
    - Версия 1.2.x Decort-SDK соответствует 3.8.5 версии платформы

## Оглавление

- [Установка](#установка)
- [Список API](#список-api)
  - [Cloudapi](#cloudapi)
  - [Cloudbroker](#cloudbroker)
- [Работа с библиотекой](#работа-с-библиотекой)
  - [Настройка конфигурации клиента](#настройка-конфигурации-клиента)
  - [Создание клиента](#создание-клиента)
  - [Создание структуры запроса](#cоздание-структуры-запроса)
  - [Выполнение запроса](#выполнение-запроса)
- [Работа с legacy клиентом](#работа-с-legacy-клиентом)
  - [Настройка конфигурации legacy клиента](#настройка-конфигурации-legacy-клиента)
  - [Создание legacy клиента](#создание-legacy-клиента)
  - [Создание структуры запроса](#cоздание-структуры-запроса)
  - [Выполнение запроса](#выполнение-запроса)

## Установка

Выполните команду в терминале:

```bash
go get -u github.com/rudecs/decort-sdk
```

## Список API

Библиотека поддерживает две группы API платформы:

- `cloudapi` - пользовательская группа, которая позволяет воспользоваться всем стардартным функционалом платформы;
- `cloudbroker` - административная группа, которая позволяет воспользоваться всем стандартным функционалом платформы и расширенными возможностями, включающими в себя управление пользователями, ресурсами, платформами размещения ресурсов и т.д.

### Cloudapi

`Cloudapi` позволяет выполнять запросы к группе пользовательских конечных точек
Данная группа ручек позволяет выполнять следующие операции в платформе:

- `Account` - управление аккаунтами - внутренними учетными записями платформы, которые являются владельцами вычислительных ресурсов;
- `BService` - управление группами виртуальных машин (computes);
- `Compute` - управление виртуальными машинами (индивидуально);
- `ComputeCI` - управление конвейром для создания виртуальных машин;
- `Disks` - управление виртуальными дисками;
- `ExtNet` - управление виртуальными сетями, отвечающими за внешний доступ;
- `FLIPgroup` - управление группами "плавающими" ip - адресами;
- `Image` - управление образами операционных систем;
- `K8CI` - управление конвейром для создания кластера;
- `K8S` - управление кластерами kubernetes;
- `KVMPPC` - создание виртуальной машины Power PC (IBM);
- `KVMx86` - создание виртуальной машины x86;
- `LB` - управление балансировщиками нагрузки;
- `Locations` - получение информации о grid площадки;
- `RG` - управление ресурсными группами аккаунта;
- `Sizes` - получение информации о потребляемых ресурсах виртуальными машинами и дисками;
- `Tasks` - получение информации о ходе выполнения асинхронных задач (например, создание кластера);
- `VINS` - управление виртуальными изолированными сетями.

### Cloudbroker

`Cloudbroker` позволяет выполнять запросы к группе пользовательских конечных точек
Данная группа ручек позволяет выполнять следующие операции в платформе:

- `Account` - управление аккаунтами - внутренними учетными записями платформы, которые являются владельцами вычислительных ресурсов;
- `Compute` - управление виртуальными машинами (индивидуально);
- `Disks` - управление виртуальными дисками;
- `ExtNet` - управление виртуальными сетями, отвечающими за внешний доступ;
- `Grid` - управление площадками;
- `Image` - управление образами операционных систем;
- `K8CI` - управление конвейром для создания кластера;
- `K8S` - управление кластерами kubernetes;
- `KVMPPC` - создание виртуальной машины Power PC (IBM);
- `KVMx86` - создание виртуальной машины x86;
- `LB` - управление балансировщиками нагрузки;
- `RG` - управление ресурсными группами аккаунта;
- `SEP` - управление storage endpoint (sep);
- `Tasks` - получение информации о ходе выполнения асинхронных задач (например, создание кластера);
- `VINS` - управление виртуальными изолированными сетями.

## Работа с библиотекой

Алгоритм работы с библиотекой выглядит следующим образом:

1. Настройка конфигурации клиента.
2. Создание клиента.
3. Создание структуры запроса.
4. Выполнение запроса.

### Настройка конфигурации клиента

Сначала, необходимо создать переменную конфигурации клиента. Конфигурация состоит как из обязательных, так и необязательных полей.
| Поле | Тип | Обязательный | Описание |
| --- | --- | --- | --- |
| AppID | string | Да | app_id ключа для выполнения запросов |
| AppSecret | string | Да | app_secret ключ для выполнения запроса |
| SSOURL | string | Да | URL адрес сервиса аутентификации и авторизации |
| DecortURL | string | Да | URL адрес платформы, с которой будет осуществляться взаимодействие |
| Retries | uint | Нет | Кол-во неудачных попыток выполнения запроса, по умолчанию - 5 |
| SSLSkipVerify | bool | Нет | Пропуск проверки подлинности сертификата, по умолчанию - true |
| Token | string | Нет | JWT токен |

Пример кода:

```go
import (
    "github.com/rudecs/decort-sdk/config"
)
func main(){
    // Настройка конфигурации
    cfg := config.Config{
        AppID:     "<APP_ID>",
        AppSecret: "<APP_SECRET>",
        SSOURL:    "https://sso.digitalenergy.online",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }
}
```

### Создание клиента

Создание клиента происходит с помощью функции-строителя `New` из основного пакета `decort-sdk`, для избежания проблем с именами, пакету можно присвоить алиас `decort`. Функция принимает конфигурацию, возвращает структуру `DecortClient`, с помощью которой можно взаимодействовать с платформой.

### Пример

```go
package main

import (
    "github.com/rudecs/decort-sdk/config"
    decort "github.com/rudecs/decort-sdk"
)

func main() {
    // Настройка конфигурации
    cfg := config.Config{
        AppID:     "<APPID>",
        AppSecret: "<APPSECRET>",
        SSOURL:    "https://sso.digitalenergy.online",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    // Создание клиента
    client := decort.New(cfg)
}
```

### Создание структуры запроса

Структуры запросов определены в пакетах:

- `pkg/cloudapi` - для `cloudapi`
- `pkg/cloudbroker` - для `cloudbroker`

В каждом пакете находятся пакеты групп API:

- **cloudapi**:
  - `pkg/cloudapi/account` - для `Account`
  - `pkg/cloudapi/bservice` - для `Basic Service`
  - `pkg/cloudapi/compute` - для `Compute`
  - `pkg/cloudapi/computeci` - для `ComputeCI`
  - `pkg/cloudapi/disks` - для `Disks`
  - `pkg/cloudapi/extnet` - для `ExtNet`
  - `pkg/cloudapi/flipgroup` - для `FLIPGroup`
  - `pkg/cloudapi/image` - для `Image`
  - `pkg/cloudapi/k8ci` - для `K8CI`
  - `pkg/cloudapi/k8s` - для `K8S`
  - `pkg/cloudapi/kvmppc` - для `KVMPPC`
  - `pkg/cloudapi/kvmx86` - для `KVMX86`
  - `pkg/cloudapi/lb` - для `LB`
  - `pkg/cloudapi/locations` - для `Locations`
  - `pkg/cloudapi/rg` - для `RG`
  - `pkg/cloudapi/sizes` - для `RG`
  - `pkg/cloudapi/tasks` - для `Tasks`
  - `pkg/cloudapi/vins` - для `VINS`
- **cloudbroker**:
  - `pkg/cloudbroker/account` - для `Account`
  - `pkg/cloudbroker/compute` - для `Compute`
  - `pkg/cloudbroker/disks` - для `Disks`
  - `pkg/cloudbroker/extnet` - для `ExtNet`
  - `pkg/cloudbroker/grid` - для `Grid`
  - `pkg/cloudbroker/image` - для `Image`
  - `pkg/cloudbroker/k8ci` - для `K8CI`
  - `pkg/cloudbroker/k8s` - для `K8S`
  - `pkg/cloudbroker/kvmppc` - для `KVMPPC`
  - `pkg/cloudbroker/kvmx86` - для `KVMX86`
  - `pkg/cloudbroker/lb` - для `LB`
  - `pkg/cloudbroker/rg` - для `RG`
  - `pkg/cloudbroker/sep` - для `SEP`
  - `pkg/cloudbroker/tasks` - для `Tasks`
  - `pkg/cloudbroker/vins` - для `VINS`

Все поля структуры имеют описание, в которых содержится:

- Их назначение;
- Обязательный или нет - поле required в комментариях;
- Доп. информация (допустимые значения, значения по умолчанию).

#### Пример комментарие структуры

```go
type CreateRequest struct {
	// ID of the resource group, which will own this VM
	// Required: true
	RGID uint64 `url:"rgId"`

	// Name of this VM.
	// Must be unique among all VMs (including those in DELETED state) in target resource group
	// Required: true
	Name string `url:"name"`

	// Number CPUs to allocate to this VM
	// Required: true
	CPU uint64 `url:"cpu"`

	// Volume of RAM in MB to allocate to this VM
	// Required: true
	RAM uint64 `url:"ram"`

	// ID of the OS image to base this VM on;
	// Could be boot disk image or CD-ROM image
	// Required: true
	ImageID uint64 `url:"imageId"`

	// Size of the boot disk in GB
	// Required: false
	BootDisk uint64 `url:"bootDisk,omitempty"`

	// ID of SEP to create boot disk on.
	// Uses images SEP ID if not set
	// Required: false
	SEPID uint64 `url:"sepId,omitempty"`

	// Pool to use if SEP ID is set, can be also empty if needed to be chosen by system
	// Required: false
	Pool string `url:"pool,omitempty"`

	// Network type
	// Should be one of:
	//	- VINS
	//	- EXTNET
	//	- NONE
	// Required: false
	NetType string `url:"netType,omitempty"`

	// Network ID for connect to,
	// for EXTNET - external network ID,
	// for VINS - VINS ID,
	// when network type is not "NONE"
	// Required: false
	NetID uint64 `url:"netId,omitempty"`

	// IP address to assign to this VM when connecting to the specified network
	// Required: false
	IPAddr string `url:"ipAddr,omitempty"`

	// Input data for cloud-init facility
	// Required: false
	Userdata string `url:"userdata,omitempty"`

	// Text description of this VM
	// Required: false
	Description string `url:"desc,omitempty"`

	// Start VM upon success
	// Required: false
	Start bool `url:"start,omitempty"`

	// Stack ID
	// Required: false
	StackID uint64 `url:"stackId,omitempty"`

	// System name
	// Required: false
	IS string `url:"IS,omitempty"`

	// Compute purpose
	// Required: false
	IPAType string `url:"ipaType,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty"`
}
```

#### Пример создания запроса для развертывания виртуальной машины:

```go
package main

import (
    "github.com/rudecs/decort-sdk/config"
    decort "github.com/rudecs/decort-sdk"
    "github.com/rudecs/decort-sdk/pkg/cloudapi/kvmx86"
)

func main() {
    // Настройка конфигурации
    cfg := config.Config{
        AppID:     "<APPID>",
        AppSecret: "<APPSECRET>",
        SSOURL:    "https://sso.digitalenergy.online",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    // Создание клиента
    client := decort.New(cfg)

    // Создание структуры запроса
    // CreateRequest - реквест на создание виртуальной машины
    req := kvmx86.CreateRequest{
        RGID:    123,
        Name:    "compute",
        CPU:     4,
        RAM:     4096,
        ImageID: 321,
    }
}
```

### Выполнение запроса

Чтобы выполнить запрос, необходимо:

1. Вызвать у клиента метод, отвечаеющий за определение группы API для взаимодействия, это может быть `.CloudAPI()`, либо `.CloudBroker()`. Данные методы возвращаеют соответствующие структуры, с помощью которых можно соверать запросы.
2. Вызвать у возвращенной структуры метод, определяющий группу ручек для взаимодействия.

   Доступные методы для `.CloudAPI()`:

   - `.Account()` - для работы с `Account`
   - `.BService()` - для работы с `BService`
   - `.Compute()` - для работы с `Compute`
   - `.ComputeCI()` - для работы с `ComputeCI`
   - `.Disks()` - для работы с `Disks`
   - `.ExtNet()` - для работы с `ExtNet`
   - `.FLIPgroup()` - для работы с `FLIPGroup`
   - `.Image()` - для работы с `Image`
   - `.K8CI()` - для работы с `K8CI`
   - `.K8S()` - для работы с `K8S`
   - `.KVMPPC()` - для работы с `KVMPPC`
   - `.KVMx86()` - для работы с `KVMX86`
   - `.LB()` - для работы с `LB`
   - `.Locations()` - для работы с `Locations`
   - `.RG()` - для работы с `RG`
   - `.Sizes()` - для работы с `Sizes`
   - `.Tasks()` - для работы с `Tasks`
   - `.VINS()` - для работы с `VINS`

   Доступные методы для `.CloudBroker()`:

   - `.Account()` - для работы с `Account`
   - `.Compute()` - для работы с `Compute`
   - `.Disks()` - для работы с `Disks`
   - `.ExtNet()` - для работы с `ExtNet`
   - `.Grid()` - для работы с `Grid`
   - `.Image()` - для работы с `Image`
   - `.K8CI()` - для работы с `K8CI`
   - `.K8S()` - для работы с `K8S`
   - `.KVMPPC()` - для работы с `KVMPPC`
   - `.KVMx86()` - для работы с `KVMX86`
   - `.LB()` - для работы с `LB`
   - `.RG()` - для работы с `RG`
   - `.SEP()` - для работы с `SEP`
   - `.Tasks()` - для работы с `Tasks`
   - `.VINS()` - для работы с `VINS`

3. Вызвать метод, отвечающий за выполнение запроса и передать в него:

- контекст;
- структуру запроса.
  У кождой группы ручек API имеются свои доступные методы, которые определяются платформой.

4. Обработать результат и ошибки.

Т.к. все вызовы методов идут последовательно, можно их объеденить в конвейер:
Общий вид вонвейра будет выглядеть так:

```go
  client.<API>.<группа>.<метод>
```

#### Пример выполнения запроса

```go
package main

import (
    "log"
    "fmt"

    "github.com/rudecs/decort-sdk/config"
    decort "github.com/rudecs/decort-sdk"
    "github.com/rudecs/decort-sdk/pkg/cloudapi/kvmx86"
)

func main() {
    // Настройка конфигурации
    cfg := config.Config{
        AppID:     "<APPID>",
        AppSecret: "<APPSECRET>",
        SSOURL:    "https://sso.digitalenergy.online",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    // Создание клиента
    client := decort.New(cfg)

    // Создание структуры запроса
    // CreateRequest - реквест на создание виртуальной машины
    req := kvmx86.CreateRequest{
        RGID:    123,
        Name:    "compute",
        CPU:     4,
        RAM:     4096,
        ImageID: 321,
    }

    //Выполнение запроса с помощью конвейера
    res, err := client.CloudAPI().KVMX86().Create(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res)
}
```

## Работа с legacy клиентом

Работа с legacy клиентом применяется для пользователей, которые не используют для авторизации decs3o.

### Настройка конфигурации legacy клиента

Сначала, необходимо создать переменную конфигурации клиента. Конфигурация состоит как из обязательных, так и необязательных полей.

| Поле          | Тип    | Обязательный | Описание                                                           |
| ------------- | ------ | ------------ | ------------------------------------------------------------------ |
| Username      | string | Да           | username legacy пользователя                                       |
| Password      | string | Да           | пароль legacy пользователя                                         |
| DecortURL     | string | Да           | URL адрес платформы, с которой будет осуществляться взаимодействие |
| Retries       | uint   | Нет          | Кол-во неудачных попыток выполнения запроса, по умолчанию - 5      |
| SSLSkipVerify | bool   | Нет          | Пропуск проверки подлинности сертификата, по умолчанию - true      |
| Token         | string | Нет          | JWT токен                                                          |

Пример кода:

```go
import (
    "github.com/rudecs/decort-sdk/config"
)
func main(){
    // Настройка конфигурации
    legacyCfg := config.LegacyConfig{
        Username:     "<USERNAME>",
        Password: "<PASSWORD>",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }
}
```

### Создание legacy клиента

Создание клиента происходит с помощью функции-строителя `NewLegacy` из основного пакета `decort-sdk`, для избежания проблем с именами, пакету можно присвоить алиас `decort`. Функция принимает конфигурацию, возвращает структуру `DecortClient`, с помощью которой можно взаимодействовать с платформой.

### Пример

```go
package main

import (
    "github.com/rudecs/decort-sdk/config"
    decort "github.com/rudecs/decort-sdk"
)

func main() {
    // Настройка конфигурации
    legacyCfg := config.LegacyConfig{
        Username:  "<USERNAME>",
        Password:  "<PASSWORD>",
        DecortURL: "https://mr4.digitalenergy.online",
        Retries:   5,
    }

    // Создание клиента
    legacyClient := decort.NewLegacy(cfg)
}
```
