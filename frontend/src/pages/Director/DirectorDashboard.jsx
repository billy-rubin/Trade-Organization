import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import InventoryReport from './InventoryReport';
import ProfitabilityReport from './ProfitabilityReport';
import TurnoverReport from './TurnoverReport';
import StoreManagement from './StoreManagement';
import EfficiencyReport from './EfficiencyReport';
import StoreEfficiency from './StoreEfficiency'; // Наш импорт 7
import ProductCustomers from './ProductCustomers'; // Наш импорт 10

const DirectorDashboard = () => {
    const [activeTab, setActiveTab] = useState('inventory');
    const navigate = useNavigate();

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('role');
        navigate('/login');
    };

    const tabs = [
        { id: 'inventory', label: 'Инвентаризация' },
        { id: 'profitability', label: 'Рентабельность' },
        { id: 'turnover', label: 'Товарооборот' },
        { id: 'stores', label: 'Управление точками' },
        { id: 'efficiency', label: 'Эффективность (Запрос 7)' },
        { id: 'prod_customers', label: 'Покупатели товара (Запрос 10)' },
        { id: 'stores', label: 'Управление точками' },
    ];

    return (
        <div style={{ padding: '20px', maxWidth: '1000px', margin: '0 auto' }}>
            <header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', borderBottom: '2px solid #333', paddingBottom: '10px', marginBottom: '20px' }}>
                <h1 style={{ fontSize: '24px' }}>Панель управления Руководителя</h1>
                <button onClick={handleLogout} style={{ background: '#ff4d4f', color: 'white', border: 'none', padding: '8px 16px', borderRadius: '4px', cursor: 'pointer' }}>
                    Выйти
                </button>
            </header>

            <nav style={{ marginBottom: '25px', display: 'flex', gap: '5px', flexWrap: 'wrap' }}>
                {tabs.map(tab => (
                    <button
                        key={tab.id}
                        onClick={() => setActiveTab(tab.id)}
                        style={{
                            padding: '10px 15px',
                            background: activeTab === tab.id ? '#001529' : '#f0f2f5',
                            color: activeTab === tab.id ? 'white' : 'black',
                            border: '1px solid #d9d9d9',
                            borderRadius: '4px 4px 0 0',
                            cursor: 'pointer'
                        }}
                    >
                        {tab.label}
                    </button>
                ))}
            </nav>

            <main style={{ background: '#fff', padding: '20px', border: '1px solid #d9d9d9', borderRadius: '4px' }}>
                {activeTab === 'inventory' && <InventoryReport />}
                {activeTab === 'profitability' && <ProfitabilityReport />}
                {activeTab === 'turnover' && <TurnoverReport />}
                {activeTab === 'efficiency' && <StoreEfficiency />}
                {activeTab === 'prod_customers' && <ProductCustomers />}
                {activeTab === 'stores' && <StoreManagement />}
            </main>
        </div>
    );
};

export default DirectorDashboard;