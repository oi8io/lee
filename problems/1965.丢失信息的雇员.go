
package problems

/**
select employee_id
from (select e.employee_id
      from Employees e
               left join Salaries s on e.employee_id = s.employee_id
      where s.employee_id is null
      union
      select s.employee_id
      from Employees e
               right join Salaries s on e.employee_id = s.employee_id
      where e.employee_id is null) t
order by employee_id;



select employee_id
from (select employee_id
      from Employees
      union all
      select employee_id
      from Salaries
      where employee_id) a
group by employee_id
having count(employee_id) = 1
order by employee_id;

**/