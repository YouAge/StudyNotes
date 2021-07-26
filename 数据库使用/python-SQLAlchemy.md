
> SQLAlchemy  Python的orm数据库操作包， 
> `pip install SQLAlchemy`
> 用法和sql语句差别多， 不过还是有一些细微差距。 sql语句才是重点，直接套就可以了


## 1.初始化连接

```python
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from app.setting.config import settings
SQLALCHEMY_DATABASE_URL = f"{settings.EXAM_SQL_TYPE}://{settings.EXAM_PGSQL_USER}:{settings.EXAM_PGSQL_PASSWORD}@{settings.EXAM_PGSQL_HOST}:{settings.EXAM_PGSQL_PORT}/{settings.EXAM_PGSQL_DB}"

engine = create_engine(
    SQLALCHEMY_DATABASE_URL
)

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

Base = declarative_base(engine)
def get_db() -> Generator:
    # 连接数据库，开关
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
```

## 创建数据模型,

```python
class UsersModel(Base):
    """考员信息模型"""
    __tablename__ = 'users_exam' # 对应数据库
    __table_args__ = ({"schema": "Kappa", 'comment': '考试人员表'}) # schema 对应pgsql中的模式
    id = Column(Integer, primary_key=True, index=True, autoincrement=True)
    user_id = Column(String(128), nullable=True, unique=True, comment='工号,唯一账号', name='userID')  # unique=True,  唯一
    name = Column(String(20), comment='姓名', nullable=False, )
    sex = Column(String(20), comment='性别', default='保密', name='gender')
    password = Column(String(128), nullable=False, comment='密码')
    is_delete = Column(Boolean, comment='逻辑删除', default=False)
    creator_time = Column(DateTime, default=datetime.now, comment='创建时间')


## 创建外键，关联
class TopicModel(Base)
    __tablename__ = 'topic'
    __table_args__ = ({"schema": "Kappa", 'comment': '题'})
    user_id = Column(String(128), nullable=True, unique=True, comment='工号,唯一账号', name='userID') 
    sn = Column(String(128), unique=True, comment='题的唯一一标识，扫码')
    is_delete = Column(Boolean, comment='逻辑删除', default=False)

    answer_id = Column(Integer, ForeignKey('Kappa.answer.id'))
    answer = relationship('AnswerModel', cascade="save-update", backref='user_answers') #


class AnswerModel(Base)
    __tablename__ = 'answer'
    __table_args__ = ({"schema": "Kappa", 'comment': '答案'})
    id = Column(Integer, primary_key=True, index=True, autoincrement=True) 
    answer = Column(String, comment='正确的选项')

    topic = relationship('TopicModel', cascade="save-update") # 反向查询
```


## 执行sql语句

```python
sql = """ SELECT * FROM "Kappa".users_exam """
data = db.execute(sql_tp).fetchall()
print(data)
```


## 基础查询

```python
db.query(UsersModel).all() # 所有[]
db.query(UsersModel.name).filter(UsersModel.sex='男').all()
db.query(UsersModel).filter(UsersModel.sex='男').first() # 返回查询的第一个

#　filter　条件， filter后面可以一直接filter
db.query(UsersModel).filter().filter().filter()

```

## 多条件查询 and_、or_

```python
from sqlalchemy import and_, or_,
db.query(UsersModel).filter(and_(UsersModel.sex='男',UsersModel.is_delete == False))
db.query(UsersModel).filter(or(UsersModel.sex='男',UsersModel.name == 'user'))
```


##　聚合查询 group_by,排序 order_by

```python
#　聚合查询
self.db.query(PaperModel.special, PaperModel.topic_type, PaperModel.work_station). \
            filter(and_(PaperModel.plant == self.plant, PaperModel.exam_status == True)) \
            .group_by(PaperModel.special, PaperModel.topic_type, PaperModel.work_station).all()
#排序
papers = db.query(PaperModel).filter(PaperModel.exam_status == True).order_by(PaperModel.paper_id.desc()).all()
```

##　关联查询，json，outerjoin

```Python
query = db.query(ExamScoreModel).filter(ExamScoreModel.is_delete == False)
query = query.join(UsersModel,UsersModel.user_id == func.cast(ExamScoreModel.user_id,String))
```

## exists、~exists 查询，subquery子查询，

>`func.cast()` 类型转换

```python
from sqlalchemy import and_, or_, func, String, exists
# exists、 ~exists取反
query = query.filter(~exists().where(and_(UsersModel.user_id == func.cast(ExamScoreModel.user_id,String),UsersModel.plant==self.plant)))
query = query.filter(exists().where(and_(UsersModel.user_id == func.cast(ExamScoreModel.user_id, String),
                         UsersModel.plant == self.plant)))

# 子查询
stmt = db.query(Address.user_id, func.count('*').label("address_count")).group_by(Address.user_id).subquery()    # 子表
print(db.query(User, stmt.c.address_count).outerjoin((stmt, User.id == stmt.c.user_id)).order_by(User.id).all()) 

```
