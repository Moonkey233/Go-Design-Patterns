#ifndef ADAPTER_H
#define ADAPTER_H
#include <iostream>

namespace adapter {
    class Device {
    public:
        Device() = default;
        virtual ~Device() = default;
        virtual void USB(const std::string &msg) = 0;
        virtual std::string Name() = 0;
    };

    class Windows final : public Device {
    public:
        Windows() = default;
        ~Windows() override = default;
        void USB(const std::string &msg) override;
        std::string Name() override { return "Windows"; }
    };

    class Phone {
    public:
        Phone() = default;
        ~Phone() = default;
        void TypeC(const std::string& msg);

    private:
        bool m_boot = false;
    };

    class OTG final : public Device {
    public:
        explicit OTG(Phone *android);
        ~OTG() override = default;
        void USB(const std::string &msg) override;
        std::string Name() override { return "Phone"; }
    private:
        Phone *m_android = nullptr;
    };

    class Client {
    public:
        Client() = default;
        ~Client() = default;
        void InsertUSB(Device *dev, const std::string &msg);
    private:
        int m_count = 0;
    };

    void TestAdapter();
}

#endif //ADAPTER_H
