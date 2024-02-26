import { configureStore } from '@reduxjs/toolkit'
import employeeReducer from '../features/employee/employeeSlice'
import { setupListeners } from '@reduxjs/toolkit/query'
import { employeeApi } from '../services/employee'

export const store = configureStore({
  reducer: {
    value: employeeReducer,
    [employeeApi.reducerPath]: employeeApi.reducer
  },
   // Adding the api middleware enables caching, invalidation, polling,
  // and other useful features of `rtk-query`.
  middleware: (getDefaultMiddleware) =>
  getDefaultMiddleware().concat(employeeApi.middleware),
})

setupListeners(store.dispatch)

/* This creates a Redux store, and also automatically configure the 
Redux DevTools extension so that you can inspect 
the store while developing.*/