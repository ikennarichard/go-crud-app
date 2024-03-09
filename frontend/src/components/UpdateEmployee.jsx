import { useState } from "react";
import { useUpdateEmployeeMutation } from "../services/apiSlice";

export default function UpdateEmployee({initial}) {
  const [inputs, setInputs] = useState({});
  const [ updateEmployee ] = useUpdateEmployeeMutation;

  function handleSubmit() {
    updateEmployee(inputs); 
  }

  return (
    <div>
      <form>
        <div>
          <label>
            Name:
            <input
             type="text" 
             name='name'
             id='name'
             onChange={(e) => setInputs(e.target.value)}

             value={initial ? initial : inputs.name ? 
              inputs.name : ''}
            />
          </label>
        </div>
      </form>
    </div>
  )
}
