#include "adapter.h"

namespace adapter {
    void Windows::USB(const std::string& msg) {
        std::cout << "Recv Windows msg: " << msg << " " << std::endl;
    }

    void Phone::TypeC(const std::string& msg) {
        std::cout << "Recv Phone msg: " << msg << " " << std::endl;
        m_boot = true;
    }

    OTG::OTG(Phone *android) {
        m_android = android;
    }

    void OTG::USB(const std::string &msg) {
        std::string phoneAdapter = msg;
        phoneAdapter = phoneAdapter.replace(msg.find("TypeC"), std::string("TypeC").length(), "USB");
        std::cout << "OTG is converting TypeC into USB" << std::endl;
        m_android->TypeC(phoneAdapter);
    }

    void Client::InsertUSB(Device *dev, const std::string &msg) {
        std::cout << dev->Name() << " Send msg: " << msg << std::endl;
        ++m_count;
        dev->USB(msg);
    }

    void TestAdapter() {
        std::cout << "---------- TestAdapter ----------" << std::endl;
        auto *windows = new Windows();
        auto *phone = new Phone();
        const auto otg = new OTG(phone);
        auto *client = new Client();
        client->InsertUSB(windows, "USB: Hello World!");
        std::cout << std::endl;
        client->InsertUSB(otg, "TypeC: Hello World!");
        std::cout << std::endl;
    }
}
