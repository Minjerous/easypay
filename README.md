# easypal


## 接口测试
**请查看http://106.55.225.88:8020/help**

**实现功能:**
+ 注册
+ 登入
+ 改密
+ 查询自己交易记录
+ 查询与特定对象的交易记录
+ 充值
+ 转账<br>

**使用技术：**
+ 预处理
+ 加盐加密码
+ mysql定时器(定时扣费)
+ jwt



```
定时器设计如下
delimiter $$
create procedure user_money()
begin
update  update  user  set  money=money-0.01 where id>0;
INSERT INTO record(pid, txt, recordtime) values( -1,系统扣费0.01, now());
end$$

call user_procedure();    --执行一次事件


DROP EVENT IF EXISTS user_event ;	
CREATE EVENT `user_event`	
ON SCHEDULE EVERY 1 DAY STARTS DATE_ADD(DATE_ADD(CURDATE(), INTERVAL 1 DAY), INTERVAL 1 HOUR) -- 每隔一天执行一次，开始执行时间为明天凌晨1点整
ON COMPLETION NOT PRESERVE
ENABLE
DO call user_money(); 

```





