//** This object is used to validate the body data */
export default {
  type: "object",
  properties: {
    url: { type: 'string' },
  },
  required: ['url']
} as const;
