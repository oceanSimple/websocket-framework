package class

import "github.com/oceanSimple/websocket-framework"

type SendMethod func(client *wsFramework.Client, data any) error
