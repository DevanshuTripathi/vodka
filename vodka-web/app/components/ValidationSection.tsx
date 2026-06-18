export default function ValidationSection() {
  return (
    <section
      id="validation"
      className="border border-slate-300 rounded-xl p-8 my-8 bg-white shadow-sm"
    >
      <h2 className="text-2xl font-bold mb-6">Validation</h2>

      <p className="text-slate-600 mb-6">
        Validate incoming request data using struct tags before processing.
      </p>

      <div className="bg-slate-950 text-white rounded-lg p-4 overflow-x-auto">
        <pre>{`type CreateUserRequest struct {
  Name  string \`validate:"required"\`
  Email string \`validate:"required,email"\`
}

if err := c.Validate(&req); err != nil {
  return err
}`}</pre>
      </div>
    </section>
  );
}
