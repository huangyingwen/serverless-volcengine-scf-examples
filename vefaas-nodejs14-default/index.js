exports.handler = async function handler (event, context) {
  return {
    statusCode: 200,
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify({ 'message': 'hello world'}),
  };
};
