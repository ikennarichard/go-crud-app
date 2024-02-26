import { Link, Params } from 'react-router-dom'



export default function Employee() {
  const id = null;
  return (
    <div>
      <p>Name</p>
      <p>Email</p>

      <div>
        <Link to={`.employee/${id}`}>Update Employee</Link>
      </div>
    </div>
  )
}
