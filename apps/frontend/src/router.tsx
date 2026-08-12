import "./index.css";
import { AppLayout } from "@/components/layouts/app-layout";
import { AuthLayout } from "@/components/layouts/auth-layout";
import { ProctectedRoute } from "@/components/protected-route";
import { PublicRoute } from "@/components/public-route";
import { CategoriesPage } from "@/pages/categories-page";
import { DashboardPage } from "@/pages/dashboard-page";
import { LandingPage } from "@/pages/landing-page";
import { SettingsPage } from "@/pages/settings-page";
import { TodosPage } from "@/pages/todos-page";
import {
  Route,
  createBrowserRouter,
  createRoutesFromElements,
} from "react-router-dom";


const routes = createRoutesFromElements(
    <>
        <Route 
            path="/"
            element={
                <PublicRoute>
                    <LandingPage />
                </PublicRoute>
            }
        />
        <Route 
            path="/auth/*"
            element={
                <PublicRoute>
                    <AuthLayout/>
                </PublicRoute>
            }
        />
        <Route 
            path="/dashboard"
            element={
                <ProctectedRoute>
                    <AppLayout>
                        <DashboardPage/>
                    </AppLayout>
                </ProctectedRoute>
            }
        />
        <Route 
            path="/todos"
            element={
                <ProctectedRoute>
                    <AppLayout>
                        <TodosPage />
                    </AppLayout>
                </ProctectedRoute>
            }
        />
        <Route 
            path="/categories"
            element={
                <ProctectedRoute>
                    <AppLayout>
                        <CategoriesPage />
                    </AppLayout>
                </ProctectedRoute>
            }
        />
        <Route 
            path="/settings"
            element={
                <ProctectedRoute>
                    <AppLayout>
                        <SettingsPage />
                    </AppLayout>
                </ProctectedRoute>
            }
        />
    </>,
);

const router = createBrowserRouter(routes);
export default router;