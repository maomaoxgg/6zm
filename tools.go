package common

import (
	"encoding/base64"
	"strings"
	"github.com/wenzhenxi/phalgo"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"strconv"
	"math/rand"
        "log"
	"io"
	"time"
	"os"
	"regexp"
	"math"
	"encoding/json"
	"reflect"
)
/*
通信接收解密
先base64解密，再des解密
 */
func Base64Des(hash string) string {
	new_hash := strings.Replace(hash, "*", "/", -1)
	jsonbyte_crypted, _ := base64.StdEncoding.DecodeString(new_hash)
	password := phalgo.Des{}

	//两种des解密，与前端沟通前慎换
	jsonbyte_data,_ := password.DesDecryptECB(jsonbyte_crypted,params.GetKey())                       //des
	//jsonbyte_data, _ := password.TripleDesDecrypt(jsonbyte_crypted, params.GetKey3(), params.GetIv())           //3des

	return string(jsonbyte_data)
}

/*
通信输出加密
先des加密，再base64加密
 */
func DesBase64(t string) string {
	orig_data := []byte(t)
	password := phalgo.Des{}

	//两种des加密，与前端沟通前慎换
	pwd_des,_ := password.DesEncryptECB(orig_data,params.GetKey(),params.GetIv())
	fmt.Println(pwd_des)//des
	//pwd_des, _ := password.TripleDesEncrypt(orig_data, params.GetKey3(), params.GetIv())                    //3des

	//pwd_des := []byte{48,231,102,105,147,85,126,187}
	hash := base64.StdEncoding.EncodeToString(pwd_des)
	new_hash := strings.Replace(hash, "/", "*", -1)
	return new_hash
}

/*
通信输出将结构体（JSON对象）先转[]byte再转string
进行DesBase64方法加密后，放进最终json包的key：hash的value中
 */
func JsonByteString(v interface{}) interface{} {

        fmt.Printf("我是返回值%+v\n",v)
	//byte_re_js, _ := json.Marshal(v)
	//
	//json_str_response :=DesBase64(string(byte_re_js))
	//
	//return json_str_response                                    //注释即加密

	return v
}

/*
根据Content-Type取request中的body值
返回解密后的json字符串
 */

func GetBodyInfo(c echo.Context) (string) {
	var (
		hash,json_str string
		result []byte
		err error
	)
	req := c.Request()
	ip := strings.Split(req.RemoteAddress(),":")[0]
	Request := phalgo.NewRequest(c)
	switch req.Header().Get("Content-Type") {
	case "application/json;charset=utf-8":
	result, err = ioutil.ReadAll(req.Body())
	if err != nil {
		log.Printf("%s\n","以下为报错↓")
		log.Print(err)
	}
		fmt.Printf("%v",string(result))
	Request.SetJson(string(result))
	hash = Request.JsonParam("hash").GetString()
	default:
	hash = Request.PostParam("hash").GetString()
	}
	switch JsonCatch[ip] {
	case 0:
		json_str = Base64Des(hash)
	case 1:
		json_str = updateCorrectStr(Request)
	}
	fmt.Println(json_str)


	return json_str
}

func GetBodyInfo1(c echo.Context) (string,*phalgo.Request) {
	var (
		hash,json_str string
		result []byte
		err error
	)
	req := c.Request()
	ip := strings.Split(req.RemoteAddress(),":")[0]
	Request := phalgo.NewRequest(c)
	switch req.Header().Get("Content-Type") {
	case "application/json;charset=utf-8":
		result, err = ioutil.ReadAll(req.Body())
		if err != nil {
			log.Printf("%s\n","以下为报错↓")
			log.Print(err)
		}
		Request.SetJson(string(result))
		hash = Request.JsonParam("hash").GetString()
	default:
		hash = Request.PostParam("hash").GetString()

	}
	switch JsonCatch[ip] {
	case 0:
		json_str = Base64Des(hash)
	case 1:
		json_str = updateCorrectStr(Request)
	}
	fmt.Println(json_str)
	Request.SetJson(json_str)
	return json_str,Request
}

func GetBodyInfoFromWeb(c echo.Context,js interface{}){
	Request := phalgo.NewRequest(c)
	RB,_ := json.Marshal(Request.Context.FormParams())
	t := reflect.TypeOf(js)
		// 进一步获取 i 的类别信息
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
			// 只有结构体可以获取其字段信息
			fmt.Printf("\n%-8v %v 个字段:\n", t, t.NumField())
			// 进一步获取 i 的字段信息
			for i := 0; i < t.NumField(); i++ {
				fmt.Println(t.Field(i).Name)
				fmt.Println(t.Field(i).Type)
				fmt.Println(t.Field(i).Type.String())
			}
		}
	json.Unmarshal(RB,js)
	fmt.Printf("我是js%+v\n",js)
}

/*
进行MD5加密
 */
func StringMd5(s string) string{
	h := md5.New()
	h.Write([]byte(s))
	pwd := hex.EncodeToString(h.Sum(nil))
	return pwd
}

/*
SHA1加密
 */
func StringSHA1(data string) string {
	t := sha1.New();
	io.WriteString(t,data);
	return fmt.Sprintf("%x",t.Sum(nil));
}


/*
生成N位数字字符串
 */
func GetRandNum(len int) string {
      var numbers string
      for i :=0;i <len;i++{
      number := strconv.Itoa(rand.Intn(10))
      numbers += number
}
	return numbers
}

/*
string转int（无error）
 */
func StringTurnInt(s string) int {
	i,_ := strconv.Atoi(s)
	return i
}

/*
[]uint8去“{”、“}”、“，”，转成[]int
 */
func ArrUint8ToInt(arr_uint8 []uint8) []int {
	var arr_int []int
	if len(arr_uint8) == 2 {                //{和}两个字节
	}else{
	arr_str := string(arr_uint8)
	arr_str = strings.Replace(arr_str,"{","",-1)
	arr_str = strings.Replace(arr_str,"}","",-1)
	arr := strings.Split(arr_str,",")
	for _,value :=range arr{
		num,_ := strconv.Atoi(value)
		arr_int=append(arr_int,num)
	      }
	}
	return arr_int
}

/*
[]uint8去“{”、“}”、“，”，转成[]int
 */
func IntToString(arr_int []int) string {
	arr_str := "{"
	for i,_ :=range arr_int{
		if i != len(arr_int)-1{
			arr_str += strconv.Itoa(arr_int[i])+","
		}else{
			arr_str +=strconv.Itoa(arr_int[i])
		}
	}
	arr_str+="}"
	fmt.Println(arr_str)
	return arr_str
}

/*
截取字符串
 */
func SubString(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

/*
time字符串换算秒数
*/
func StringTimeToSeconds(arr_uint8 []uint8) int {
	arr_time := strings.Split(string(arr_uint8),":")
	seconds := StringTurnInt(arr_time[0])*3600 +StringTurnInt(arr_time[1])*60+StringTurnInt(arr_time[2])
	return seconds
}

/*
second换算汉字天数
*/
func SecondsToCharacterTime(seconds int) string {
	var (
		s string
		day,hour,min int
	)
	if seconds >86400 {
		day = int(math.Floor(float64(seconds/86400)));hour = int(math.Floor(float64((seconds - 86400*day)/3600)));min = int(math.Floor(float64((seconds - 86400*day-3600*hour)/60)))
		s = phalgo.IntTurnString(day)+"天"+phalgo.IntTurnString(hour)+"小时"+phalgo.IntTurnString(min)+"分钟"
	}else if seconds < 86400{
		hour = int(math.Floor(float64((seconds)/3600)));min = int(math.Floor(float64((seconds -3600*hour)/60)))
		s = phalgo.IntTurnString(hour)+"小时"+phalgo.IntTurnString(min)+"分钟"
	}
	return s
}

/*
生成三方登录id
*/
func CreateThirdId() string {
	strDay := phalgo.IntTurnString(time.Now().YearDay())
	strNa := phalgo.IntTurnString(time.Now().Nanosecond())
	lenStrDay := len(strDay)
	lenStrNa := len(strNa)
	if lenStrDay != 3{
		var zero string
		for i:=0;i<3-lenStrNa;i++{
			zero+="0"
		}
		strDay = zero + strDay
	}
	if lenStrNa != 9{
		var zero string
		for i:=0;i<9-lenStrNa;i++{
			zero+="0"
		}
		strNa = zero + strNa
	}

	return GetRandNum(2)+strDay+strNa
}

func CreateLog(a ...interface{}){
	logFilename := phalgo.GetPath()+"/Runtime/"+time.Now().Format("2006-01-02")+".log"
	logFile, _ := os.OpenFile(logFilename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0777) // append data to the file when writing.
	defer logFile.Close()
	logger := log.New(logFile,"\r\n", log.Ldate | log.Ltime | log.Llongfile)
	logger.Println(a...)
}

//手机号有效性验证
const PHONE_RGX = `^1(3[0-9]|4[5,7,9]|5[0,1,2,3,5,6,7,8,9]|7[0,1,3,5,6,7,8]|8[0-9])\d{8}$`
func ValidateMobile(mobile int) bool {
	rgx := regexp.MustCompile( PHONE_RGX )
	s := strconv.Itoa(mobile)
	return rgx.MatchString(s)
}

//处理往字符串中增减字符串
//aor 1增加   2 去除
func StringAddOrRemove(datas , str string, aor int) string {
	s := StrToSlice(datas) //字符串转数组
	if aor == 2 {
		index := 0
		endIndex := len(s) - 1
		var result = make([]string, 0)
		for k, v := range s {
			if v == str {
				result = append(result, s[index:k]...)
				index = k + 1
			} else if k == endIndex {
				result = append(result, s[index:endIndex+1]...)
			}
		}
		cc := strings.Join(result,",")
		return cc
	}else{
		//先判断之前是否存在
		for _, val := range s {
			if val == str {
				return datas
			}
		}
		bbb := append(s,str) //新增数组元素
		cc := strings.Join(bbb,",") //数组转为字符串，`，`号分割
		return cc
	}

}
func StrToSlice(str string) []string {
	canSplit := func (c rune)  bool { return c == ','}
	return strings.FieldsFunc(str,canSplit)  //字符串转数组
}

//字符串数组转int
func StrArrayToInt(arr []string) []int {
	var arr_int []int
	for _,value :=range arr{
		num,_ := strconv.Atoi(value)
		arr_int=append(arr_int,num)
	}
	return  arr_int
}
//int数组转字符串
func IntArrayToString(arr []int) []string {
	var arr_string []string
	for _,value :=range arr{
		num := strconv.Itoa(value)
		arr_string=append(arr_string,num)
	}
	return  arr_string
}

func updateCorrectStr(r *phalgo.Request) string{
	RB,_ := json.Marshal(r.Context.FormParams())
	RBReplace := strings.Replace(string(RB),"],\"","]|\"",-1)
	RBArray := strings.Split(RBReplace,"|")
	var(
		newArray []string
		newValue string
	)
	for _,v :=range RBArray{
		if (strings.Contains(v,"%i") || strings.Contains(v,"%a")) {
			newValue = strings.Replace(v,"[\"","",-1);newValue = strings.Replace(newValue,"\"]","",-1);
			newValue = strings.Replace(newValue,"%i","",-1)
			newValue = strings.Replace(newValue,"%a","",-1)
		}else{
			newValue = strings.Replace(v,"[\"","\"",-1);newValue = strings.Replace(newValue,"\"]","\"",-1);
		}
		newArray = append(newArray,newValue)
	}
	RBFinal := strings.Join(newArray,",")
	fmt.Printf("我是json_str+%s\n",RBFinal)
	return RBFinal
}
