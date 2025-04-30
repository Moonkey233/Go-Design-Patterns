#ifndef PROXY_H
#define PROXY_H
#include <iostream>
#include <string>

namespace proxy {
    class Server {
    public:
        virtual ~Server() = default;
        virtual std::pair<int, std::string> request(std::string, std::string) const = 0;
    };

    class Application final : public Server {
    public:
        std::pair<int, std::string> request(std::string, std::string) const override;
    };

    class Nginx final : public Server {
    public:
        explicit Nginx(Application *app);
        static bool checkAccess(const std::string& url, const std::string& method) ;
        std::pair<int, std::string> request(std::string, std::string) const override;
    private:
        Application *m_app = nullptr;
    };

    void TestProxy();
}

#endif //PROXY_H
