# dotsql


user
----
office.name = $l

1. attach table to each field

{table=location, column=office, alias="a", pk="id"}.{leaf=name}

2.  join <table>  <alias> on <alias>.<pk> = ?.<column>
    join location a       on a.id =         user.office

3. select <alias>.<leaf> 

    select a.name


user
----
office.manager.name = $l

1. attach table to each field

{table=user, alias="user"}.{table=location, column=office, alias="a", pk="id"}.{table=user, column=manager, alias="b", pk="id"}.{leaf=name}

2.  join <table>  <alias> on <alias>.<pk> = <parent.alias>.<column>
    join location a       on a.id =         user.office
    join user     b       on b.id =         a.manager

3. select <parent.alias>.<leaf>

   select a.name