# crash

## login
```
mysql -u root -p 12345678
```

## Create database
```
create database sqlbzbh;
```

```
show databases;
show tables;
desc [TABLE NAME];
```

```
use sqlbzbh;
```

## alter table
```
alter table customers drop column cust_zip;
alter table customers add cust_zip char(10) not null;
alter table customers change cust_zip cust_modify char(10) not null;
alter table customers modify column cust_name var(255) NOT NULL;
```

## insert record
```
insert into customers(cust_id, cust_name, cust_address, cust_city, cust_state, cust_zip, cust_country, cust_contact)
values(10006, 'mao', 'renmin rd 1', 'beijing', 'BJ', 10001, 'PRC', "zhou");
```

## delete record
```
delete from customers where cust_id=10006;
```

## update record
```
update customers set cust_zip='44444' where cust_id=10001;
```


