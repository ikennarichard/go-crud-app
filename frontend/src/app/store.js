import { configureStore } from '@reduxjs/toolkit'
import { apiSlice } from '../services/apiSlice'

export const store = configureStore({
  reducer: {
    [apiSlice.reducerPath]: apiSlice.reducer
  },
  
  middleware: (getDefaultMiddleware) =>
  getDefaultMiddleware().concat(apiSlice.middleware),
})

/* This creates a Redux store, and also automatically configure the 
Redux DevTools extension so that you can inspect 
the store while developing.*/