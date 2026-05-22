import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/Login';
import ProtectedRoute from './components/ProtectedRoute';

// Заглушки для страниц
import SellerDashboard from './pages/Seller/SellerDashboard';
import ManagerDashboard from './pages/Manager/ManagerDashboard';
import DirectorDashboard from './pages/Director/DirectorDashboard';

function App() {
  return (
      <Router>
        <Routes>
          {/* Публичный роут */}
          <Route path="/login" element={<Login />} />

          {/* Роуты Продавца */}
          <Route path="/trade/*" element={
            <ProtectedRoute allowedRoles={['role_seller', 'role_director']}>
              <SellerDashboard />
            </ProtectedRoute>
          } />

          {/* Роуты Менеджера по закупкам */}
          <Route path="/supply/*" element={
            <ProtectedRoute allowedRoles={['role_purchase_manager', 'role_director']}>
              <ManagerDashboard />
            </ProtectedRoute>
          } />

          {/* Роуты Руководителя */}
          <Route path="/reports/*" element={
            <ProtectedRoute allowedRoles={['role_director']}>
              <DirectorDashboard />
            </ProtectedRoute>
          } />

          {/* Редирект по умолчанию */}
          <Route path="*" element={<Navigate to="/login" replace />} />
        </Routes>
      </Router>
  );
}

export default App;