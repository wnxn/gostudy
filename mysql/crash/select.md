# Select

> https://www.cnblogs.com/xiaobingqianrui/p/10020059.html

```
select -- DISTINCT -- from --  where -- GROUP BY -- HAVING -- ORDER BY -- LIMIT
```

## distinct
```
select distinct vend_id from products;
+---------+
| vend_id |
+---------+
|    1001 |
|    1002 |
|    1003 |
|    1005 |
+---------+
4 rows in set (0.00 sec)
```

## group by
```
select count(*) from products where vend_id =1003;
+----------+
| count(*) |
+----------+
|        7 |
+----------+
1 row in set (0.00 sec)
```

```
select vend_id, count(*) as num_prods from products group by vend_id order by  num_prods desc ;
+---------+-----------+
| vend_id | num_prods |
+---------+-----------+
|    1003 |         7 |
|    1001 |         3 |
|    1002 |         2 |
|    1005 |         2 |
+---------+-----------+
4 rows in set (0.00 sec)
```

## having

```
select vend_id, count(*) as num_prods from products group by vend_id order by vend_id desc limit 1,3;
+---------+-----------+
| vend_id | num_prods |
+---------+-----------+
|    1003 |         7 |
|    1002 |         2 |
|    1001 |         3 |
+---------+-----------+
3 rows in set (0.00 sec)
```