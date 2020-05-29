select employees.emp_no, departments.dept_name, titles.title, employees.first_name, employees.last_name,
 employees.hire_date, year(now())-year(employees.hire_date) as working_for
 from employees inner join dept_emp on employees.emp_no = dept_emp.emp_no
 inner join departments on dept_emp.dept_no = departments.dept_no inner join titles on
 employees.emp_no = titles.emp_no
 where titles.to_date > now() and month(employees.hire_date) = month(now()) and mod(year(now())-year(employees.hire_date), 5) = 0