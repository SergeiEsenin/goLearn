 select departments.dept_no, count(*) as emp_count, sum(salaries.salary) from departments inner join dept_emp
 on departments.dept_no = dept_emp.dept_no
 inner join employees on dept_emp.emp_no = employees.emp_no inner join salaries on employees.emp_no = salaries.emp_no
 where salaries.to_date > now() and
 dept_emp.to_date > now() group by departments.dept_name order by departments.dept_no;