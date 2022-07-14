//表: Employee
//
//
//+--------------+---------+
//| Column Name  | Type    |
//+--------------+---------+
//| id           | int     |
//| name         | varchar |
//| salary       | int     |
//| departmentId | int     |
//+--------------+---------+
//Id是该表的主键列。
//departmentId是Department表中ID的外键。
//该表的每一行都表示员工的ID、姓名和工资。它还包含了他们部门的ID。
//
//
//
//
// 表: Department
//
//
//+-------------+---------+
//| Column Name | Type    |
//+-------------+---------+
//| id          | int     |
//| name        | varchar |
//+-------------+---------+
//Id是该表的主键列。
//该表的每一行表示部门ID和部门名。
//
//
//
//
// 公司的主管们感兴趣的是公司每个部门中谁赚的钱最多。一个部门的 高收入者 是指一个员工的工资在该部门的 不同 工资中 排名前三 。
//
// 编写一个SQL查询，找出每个部门中 收入高的员工 。
//
// 以 任意顺序 返回结果表。
//
// 查询结果格式如下所示。
//
//
//
// 示例 1:
//
//
//输入:
//Employee 表:
//+----+-------+--------+--------------+
//| id | name  | salary | departmentId |
//+----+-------+--------+--------------+
//| 1  | Joe   | 85000  | 1            |
//| 2  | Henry | 80000  | 2            |
//| 3  | Sam   | 60000  | 2            |
//| 4  | Max   | 90000  | 1            |
//| 5  | Janet | 69000  | 1            |
//| 6  | Randy | 85000  | 1            |
//| 7  | Will  | 70000  | 1            |
//+----+-------+--------+--------------+
//Department  表:
//+----+-------+
//| id | name  |
//+----+-------+
//| 1  | IT    |
//| 2  | Sales |
//+----+-------+
//输出:
//+------------+----------+--------+
//| Department | Employee | Salary |
//+------------+----------+--------+
//| IT         | Max      | 90000  |
//| IT         | Joe      | 85000  |
//| IT         | Randy    | 85000  |
//| IT         | Will     | 70000  |
//| Sales      | Henry    | 80000  |
//| Sales      | Sam      | 60000  |
//+------------+----------+--------+
//解释:
//在IT部门:
//- Max的工资最高
//- 兰迪和乔都赚取第二高的独特的薪水
//- 威尔的薪水是第三高的
//
//在销售部:
//- 亨利的工资最高
//- 山姆的薪水第二高
//- 没有第三高的工资，因为只有两名员工
// 👍 630 👎 0

package cn

//There is no code of Go type for this problem

/*
SELECT d.Name  AS 'Department',
e1.Name AS 'Employee',
e1.Salary
FROM Employee e1
JOIN
Department d ON e1.DepartmentId = d.Id
WHERE 3 > (SELECT COUNT(DISTINCT e2.Salary)
FROM Employee e2
WHERE e2.Salary > e1.Salary
AND e1.DepartmentId = e2.DepartmentId);
*/
