import "./index.css";
import router from "./router.tsx";
import { ThemeProvider } from "@/components/theme-provider";
import { Toaster } from "@/components/ui/sonner";
import { CLERK_PUBLISHER_KEY } from "@/config/env";
import { ClerkProvider } from "@clerk/clerk-react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { RouterProvider } from "react-router-dom";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 5 * 60 * 1000, // 5 minutes
      gcTime: 10 * 60 * 1000, // 10 minutes
    },
  },
});

if(!CLERK_PUBLISHER_KEY){
  throw new Error("Missing Publishable Key");
}

const clerkPubKey = CLERK_PUBLISHER_KEY;

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ClerkProvider publishableKey={clerkPubKey}>
      <QueryClientProvider client={queryClient}>
        <ThemeProvider defaultTheme="system" storageKey="taskForge-ui-theme">
          <RouterProvider router={router} />
          <Toaster />
          <ReactQueryDevtools initialIsOpen={false}/>
        </ThemeProvider>
      </QueryClientProvider>
    </ClerkProvider>
  </StrictMode>
)