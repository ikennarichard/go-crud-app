import { useMemo } from "react"
import { useGetEmployeesQuery, useDeleteEmployeeMutation } from "../services/apiSlice"
import AddEmployee from "../components/AddEmployee";

export default function Home() {

  const { 
    data: employees=[], 
    error,
    isError,
    isFetching,
    isLoading,
    isSuccess 
  } = useGetEmployeesQuery()

  const [ deleteEmployee ] = useDeleteEmployeeMutation();

  const handleDelete = (id) => {
    deleteEmployee(id)
  }

  const sortedEmployees = useMemo(() => {
    const sortedEmployees = employees.slice()
    
    sortedEmployees.sort((a, b) => b.CreatedAt.localeCompare(a.CreatedAt))
    console.log(sortedEmployees)
    return sortedEmployees
  }, [employees])

  let content

  if(isLoading) {
    content = <div>Loading...</div>
  } else if (isSuccess) {
    const renderedEmployees = sortedEmployees?.map((e) => (
      <div key={e.ID}>
          <b>{e['Name']}</b>
          <p>{e['Email']}</p>
          <div>
            <button>Edit</button>
            <button onClick={() => handleDelete(e.ID)}>Delete</button>
          </div>
          <hr />
      </div>
    ))

    content = <div className={{disabled: isFetching}}>{renderedEmployees}</div>
  } else if(isError) {
    content = <div>{error.toString()}</div>
  }

  return (
    <>
      <AddEmployee/>
      <section className="employees-list">
        <h2>Employees</h2>
        {content}
      </section>
    </>
  )
}

/*
You can always check data, status, and 
error to determine the right UI to render. 
In addition, useQuery also provides utility booleans like 
isLoading, isFetching, isSuccess, and isError for the latest request.*/ 