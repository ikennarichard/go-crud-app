# Frontend

## RTK Query

By default, `query endpoints` will use a `GET HTTP request`, but you can override that by returning an object like `{url: '/posts', method: 'POST', body: newPost}` instead of just the URL string itself. You can also define several other options for the request this way, such as setting `headers`.

RTK Query's React integration will automatically generate React hooks for every endpoint we define!

### Configuring the store

The API slice generates a custom middleware that needs to be added to the store. This middleware must be added as well - it manages cache lifetimes and expiration.

### Using queries and mutations

Currently, `<PostsList>` is specifically importing `useSelector`, `useDispatch, and useEffect`, reading posts data and loading state from the store, and dispatching the fetchPosts() thunk on mount to trigger the data fetch. The `useGetPostsQueryHook` replaces all of that!

## Notes

```js
  //containerClassname
    const renderedPosts = sortedPosts.map(post => (
      <PostExcerpt key={post.id} post={post} />
    ))

    const containerClassname = classnames('posts-container', {
      disabled: isFetching
    })

    content = <div className={containerClassname}>{renderedPosts}</div>
```
