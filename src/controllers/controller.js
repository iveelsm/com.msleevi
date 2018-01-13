// @flow

export const homePage = () => null

export const helloPage = () => ({
  hello: { message: 'Server-side preloaded message' },
})

export const helloAsyncPage = () => ({
  hello: { messageAsync: 'Server-side preloaded message for asyncPage' },
})

export const helloEndpoint = (num: number) => ({
  serverMessage: `Hello fomr the server! (received ${num})`,
})
