- [MySql](#mysql)
  - [cli](#cli)
  - [join](#join)
    - [inner join](#inner-join)
    - [outer join](#outer-join)
    - [left join](#left-join)
    - [right join](#right-join)
    -
## MySql

### cli

```
# bin
/usr/local/opt/mysql@5.7/bin/mysql
```

```
show global variables like "%datadir%";
```

### join

#### inner join

- default join type is inner join
- `inner join` == `join`

```sql
select *
from store
join address
on store.address_id = address.address_id
```

#### outer join

- `outer join` == `left outer join` == `left join`
- `outer right join` == `right join`
#### left join

- get everything left side of `join` keyword


#### right join

- get everything right side of `join` keyword

```sql
select *
from store
right join address
on store.address_id = address.address_id
```
