import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import CreateRequest from './CreateRequest';
import GenerateOrder from './GenerateOrder';
import ReceiveOrder from './ReceiveOrder';
import SupplierDeliveries from './SupplierDeliveries'; // Наш импорт 6
import OrderDetailsReport from './OrderDetailsReport'; // Наш импорт 9

const ManagerDashboard = () => {
    const [activeTab, setActiveTab] = useState('request');
    const navigate = useNavigate();

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('role');
        navigate('/login');
    };

    const tabs = [
        { id: 'request', label: 'Новая заявка точки' },
        { id: 'order', label: 'Заказ поставщику' },
        { id: 'receive', label: 'Приемка товара' },
        { id: 'deliveries', label: 'Архив поставок (Запрос 6)' },
        { id: 'details', label: 'Состав заказа (Запрос 9)' },
    ];

    return (
        <div style={{ padding: '20px', maxWidth: '800px', margin: '0 auto' }}>
            <header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', borderBottom: '2px solid #eee', paddingBottom: '10px', marginBottom: '20px' }}>
                <h2>Панель Менеджера по закупкам</h2>
                <button onClick={handleLogout} style={{ background: '#ff4d4f', color: 'white', border: 'none', padding: '8px 16px', borderRadius: '4px', cursor: 'pointer' }}>
                    Выйти
                </button>
            </header>

            <div style={{ marginBottom: '20px', display: 'flex', gap: '10px' }}>
                <button
                    onClick={() => setActiveTab('request')}
                    style={{ padding: '10px 20px', background: activeTab === 'request' ? '#1890ff' : '#f0f0f0', color: activeTab === 'request' ? 'white' : 'black', border: 'none', borderRadius: '4px', cursor: 'pointer' }}
                >
                    1. Новая заявка
                </button>
                <button
                    onClick={() => setActiveTab('order')}
                    style={{ padding: '10px 20px', background: activeTab === 'order' ? '#1890ff' : '#f0f0f0', color: activeTab === 'order' ? 'white' : 'black', border: 'none', borderRadius: '4px', cursor: 'pointer' }}
                >
                    2. Заказ поставщику
                </button>
                <button
                    onClick={() => setActiveTab('receive')}
                    style={{ padding: '10px 20px', background: activeTab === 'receive' ? '#1890ff' : '#f0f0f0', color: activeTab === 'receive' ? 'white' : 'black', border: 'none', borderRadius: '4px', cursor: 'pointer' }}
                >
                    3. Приемка товара
                </button>
            </div>

            <main style={{ background: '#fff', padding: '20px', border: '1px solid #d9d9d9', borderRadius: '4px' }}>
                {activeTab === 'request' && <CreateRequest />}
                {activeTab === 'order' && <GenerateOrder />}
                {activeTab === 'receive' && <ReceiveOrder />}
                {activeTab === 'deliveries' && <SupplierDeliveries />}
                {activeTab === 'details' && <OrderDetailsReport />}
            </main>
        </div>
    );
};

export default ManagerDashboard;