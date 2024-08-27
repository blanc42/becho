export default function LoadingPage() {
  return (
    <div className="flex items-center justify-center min-h-screen bg-background">
      <div className="text-center">
        <div className="animate-spin rounded-full h-32 w-32 border-t-2 border-b-2 border-primary mx-auto mb-8"></div>
        <h1 className="text-4xl font-bold text-foreground">Loading...</h1>
        <p className="text-xl text-muted-foreground mt-4">Please wait while we prepare your content</p>
      </div>
    </div>
  );
}