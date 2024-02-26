import React from 'react'

export default function UpdateEmployee() {
  return (
    <div>
      <form>
        <div>
          <label htmlFor="name">
            Name:
            <input type="text" name='name' id='name'/>
          </label>
        </div>
        <div>
          <label htmlFor="email">
            Email:
            <input type="emal" name='email' id='enail'/>
          </label>
        </div>
      </form>
    </div>
  )
}
