# 　TODO   APSchedule 定时框架

# 　TODO　官方文档：https://apscheduler.readthedocs.io/en/latest/userguide.html#basic-concepts

# 　简介

"""
APScheduler全程为Advanced Python Scheduler，是一款轻量级的Python任务调度框架。
它允许你像Cron那样安排定期执行的任务，并且支持Python函数或任意可调用的对象。

安装：
pip install apscheduler 
方法二：如果pip不起作用，可以从pypi上下载最新的源码包(https://pypi.python.org/pypi/APScheduler/)进行安装
python setup.py install 

"""

# 　TODO　APScheduler组件


"""
triggers（触发器）： 触发器中包含调度逻辑，每个作业都由自己的触发器来决定下次运行时间。除了他们自己初始配置意外，触发器完全是无状态的。
        当你调度作业的时候，你需要为这个作业选择一个触发器，用来描述这个作业何时被触发，APScheduler有三种内置的触发器类型:
        date 一次性指定日期 
        interval 在某个时间范围内间隔多长时间执行一次 
        cron 和Linux crontab格式兼容，最为强大


job stores（作业存储器）：存储被调度的作业，默认的作业存储器只是简单地把作业保存在内存中，其他的作业存储器则是将作业保存在数据库中。
    当作业被保存到一个持久化的作业存储器中的时候，该作业的数据会被序列化，并在加载时被反序列化。作业存储器不能共享调度器。

executors（执行器）：处理作业的运行，他们通常通过在作业中提交指定的可调用对象到一个线程或者进城池来进行。当作业完成时，执行器将会通知调度器。

schedulers（调度器）：配置作业存储器和执行器可以在调度器中完成，例如添加、修改和移除作业。根据不同的应用场景可以选用不同的调度器
    以下7中调度器：
    BlockingScheduler : 当调度器是你应用中唯一要运行的东西时。 
    BackgroundScheduler : 当你没有运行任何其他框架并希望调度器在你应用的后台执行时使用（充电桩即使用此种方式）。 
    AsyncIOScheduler : 当你的程序使用了asyncio（一个异步框架）的时候使用。 
    GeventScheduler : 当你的程序使用了gevent（高性能的Python并发框架）的时候使用。 
    TornadoScheduler : 当你的程序基于Tornado（一个web框架）的时候使用。 
    TwistedScheduler : 当你的程序使用了Twisted（一个异步框架）的时候使用 
    QtScheduler : 如果你的应用是一个Qt应用的时候可以使用

"""
from apscheduler.schedulers.blocking import BlockingScheduler
from datetime import date, datetime


def test():
    print("测试函数")


# 　TODO 一次性调度器
scheduler = BlockingScheduler()
# 2019-12-12 运行一次
scheduler.add_job(test, "data", run_date=date(2019, 12, 12), args=['text'])
# 2019-12-12 12:12:12运行一次
scheduler.add_job(test, 'date', run_date=datetime(2019, 12, 12, 12, 12, 12), args=['text'])

# 　TODO 间接调度
"""
weeks (int) – 间隔几周 
days (int) – 间隔几天 
hours (int) – 间隔几小时 
minutes (int) – 间隔几分钟 
seconds (int) – 间隔多少秒 
start_date (datetime|str) – 开始日期 
end_date (datetime|str) – 结束日期 
timezone (datetime.tzinfo|str) – 时区
"""
# 2小时执行一次
scheduler.add_job(test, 'interval', hours=2)

# 　TODO 间接调度
"""
year (int|str) – 年，4位数字 
month (int|str) – 月 (范围1-12) 
day (int|str) – 日 (范围1-31) 
week (int|str) – 周 (范围1-53) 
day_of_week (int|str) – 周内第几天或者星期几 (范围0-6 或者 mon,tue,wed,thu,fri,sat,sun) 
hour (int|str) – 时 (范围0-23) 
minute (int|str) – 分 (范围0-59) 
second (int|str) – 秒 (范围0-59) 
start_date (datetime|str) – 最早开始日期(包含) 
end_date (datetime|str) – 最晚结束时间(包含) 
timezone (datetime.tzinfo|str) – 指定时区 

"""
# 每天早上8点15执行
scheduler.add_job(test, 'cron', day_of_week='*', hour='8', minute='15')
# test将会在6,7,8,11,12月的第3个周五的1,2,3点运行
scheduler.add_job(test, 'cron', month='6-8,11-12', day='3rd fri', hour='0-3')
# 截止到2016-12-30 00:00:00，每周一到周五早上五点半运行test
scheduler.add_job(test, 'cron', day_of_week='mon-fri', hour=5, minute=30, end_date='2016-12-31')
