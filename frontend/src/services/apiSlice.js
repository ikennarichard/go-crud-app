import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

export const apiSlice = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({ baseUrl: 'http://localhost:5000' }),
  tagTypes: ['Employee'],

  endpoints: (builder) => ({
    getEmployees: builder.query({
      query: () => '/employees',
      providesTags: ['Employee']
    }),
    
    deleteEmployee: builder.mutation({
      query: id => ({
        url: `/employee/${id}`,
        method: 'DELETE',
      }),
      invalidatesTags:['Employee']
    }),
    
    addEmployee: builder.mutation({
      query: employee => ({
        url: '/employee',
        method: 'POST',
        body: employee
      }),
      invalidatesTags:['Employee']
    }),

    updateEmployee: builder.mutation({
      query: employee => ({
        url: '/employee',
        method: 'PUT',
        body: employee
      }),
      invalidatesTags:['Employee']
    }),
  }),

})

export const { 
  useGetEmployeesQuery, 
  useDeleteEmployeeMutation, 
  useAddEmployeeMutation, 
  useUpdateEmployeeMutation } = apiSlice