#include <iostream>

template <class T>
struct Node {
    T data;
    Node<T>* next;
};

template <class T>
class LinkedList {
private:
    Node<T>* head;
public:
    LinkedList() {
        head = nullptr;
    }
    int size() {
        int count = 0;
        Node<T> *temp = head;
        while(temp != nullptr) {
            ++count;
            temp = temp->next;
        }
        return count;
    }
    void insert(T data, int index = 0) {
        Node<T>* newNode = new Node<T>;
        newNode->data = data;
        if (index == 0) {
            newNode->next = head;
            head = newNode;
        } else {
            Node<T>* temp = head;
            for (int i = 0; i < index - 1; i++) {
                temp = temp->next;
            }
            newNode->next = temp->next;
            temp->next = newNode;
        }
    }
    void remove(int index) {
        if (head == nullptr || index < 0 || index > this->size() - 1) {
            return;
        }
        Node<T> *temp = head;
        if (index == 0) {
            head = head->next;
            delete temp;
        } else {
            for(int i = 0; i < index - 1; i++) {
                temp = temp->next;
            }
            Node<T> *newTemp = temp->next;
            temp->next = temp->next->next;
            delete newTemp;
        }
    }
    void reverse() {
        Node<T> *prev = nullptr, *curr = head, *next = nullptr;
        while(curr != nullptr) {
            next = curr->next;
            curr->next = prev;
            prev = curr;
            curr = next;
        }
        head = prev;
    }
    void print() {
        Node<T>* temp = head;
        while(temp->next != nullptr) {
            std::cout << temp->data << " -> ";
            temp = temp->next;
        }
        std::cout << temp->data << std::endl;
    }
};

int main() {
    LinkedList<int> ll;
    return 0;
}

