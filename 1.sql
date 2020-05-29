select titles.title, employees.first_name, employees.last_name, salaries.salary from employees inner join dept_manager on employees.emp_no = dept_manager.emp_no
inner join titles on employees.emp_no = titles.emp_no inner join salaries on employees.emp_no = salaries.emp_no
where salaries.to_date > now() and titles.to_date > now() and dept_manager.to_date > now();