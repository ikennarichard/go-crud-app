import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  employees: [],
}

export const employeeSlice = createSlice({
  name: 'employee',
  initialState,
  reducers: {
    getEmployee: (state) => {
      state.value += 1
    },
  },
})

// Action creators are generated for each case reducer function
export const { getEmployee } = employeeSlice.actions

export default employeeSlice.reducer