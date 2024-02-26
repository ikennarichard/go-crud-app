import { useSelector } from "react-redux"
import { getEmployee } from "../features/employee/employeeSlice"

export default function Home() {
  const employees = useSelector((state) => state.value.employees)
  // const dispatch = useDispatch()
  return (
    <div>
      <h1>Hello Employees</h1>
      <div></div>
    </div>
  )
}
