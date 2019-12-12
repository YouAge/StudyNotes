### url_for
	url_for 第一个常数是 视图函数






{% set user=g.csms——user  %}





card['status']["retweeted_status"]["longText"]["longTextContent"]



### flask参数请求
```python

# GET请求，post请求
@app.route('/aos_started/',methods=['GET','POST'])
def aos_conter():
    # pid = request.get_data('pid')
    if request.methods == 'GET' 
	    pid = request.args.get('pid') 
	    spider_count = request.args.get('spider_count')
	    print(pid,spider_count)
	    response = make_response(spider_count, 200)
	 if request.methods = POST:
	 	request.form.get('参数名')
    return response


# #返回状态码
from flask import make_response
# make_response()，相当于DJango中的HttpResponse。
@blue.route('/makeresponse/')
def make_response_function():
    temp = render_template('hello.html')
    response = make_response(temp, 200) 
    return response



request的常用属性

args：获取GET,post请求参数 -->request.agrs.get('参数名')
# form：获取POST请求表格
files 获取上传文件
base_url 获取请求路径
host 获取ip和端口

```










