
package api

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
	 js "github.com/dop251/goja"
	"github.com/etherlabsio/go-m3u8/m3u8"
	"github.com/tidwall/gjson"
        "github.com/YanG-1989/list/blob/main/golang/douyu.go"
	"github.com/YanG-1989/list/blob/main/golang/bilibili.go"
	"github.com/YanG-1989/list/blob/main/golang/douyin.go"
	"github.com/YanG-1989/list/blob/main/golang/huya.go"
	"github.com/YanG-1989/list/blob/main/golang/youtube.go"
)
func Tool(w http.ResponseWriter, r *http.Request)  {
    // todo
}
