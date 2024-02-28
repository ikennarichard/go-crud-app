import { useState } from 'react';
import { useAddEmployeeMutation } from '../services/apiSlice';


export default function AddEmployee() {
  const [inputs, setInputs] = useState({});
  const [ addEmployee, {isLoading} ] = useAddEmployeeMutation();
  
  const handleChange = e => {
    const name = e.target.name;
    const value = e.target.value;
    setInputs(values => ({...values, [name]: value}))
  }

  function handleSubmit(e) {
    e.preventDefault();
    addEmployee(inputs); 
  }
  return (
    <div>
      <form onSubmit={handleSubmit}>
        <p>
          <label htmlFor="name">Name:
          <input 
            type="text" 
            name='name'
            value={inputs.name || ''}
            onChange={handleChange} 
          />
          </label>
        </p>
        <p>
          <label htmlFor="email">Email:
          <input 
            type="email" 
            name='email'
            value={inputs.email || ''}
            onChange={handleChange} 
          />
          </label>
        </p>
        <button disabled={isLoading} type="submit">Submit</button>
      </form>
    </div>
  )
}
