#include <iostream>

template <class T>
struct Node {
    T data;
    Node<T> *prev, *next;
};

template <class T>
class DoublyLinkedList {
private:
    Node<T> *head;
public:
    DoublyLinkedList() {
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
        Node<T>* newNode = new Node<T>();
        newNode->data = data;
        if(index == 0) {
            if (head != nullptr) {
                head->prev = newNode;
            }
            newNode->prev = nullptr;
            newNode->next = head;
            head = newNode;
        } else {
            Node<T>* temp = head;
            for (int i = 0; i < index - 1; i++) {
                temp = temp->next;
            }
            newNode->prev = temp;
            newNode->next = temp->next;
            temp->next->prev = newNode;
            temp->next = newNode;
        }
    }
    void remove(int index) {
        if(head == nullptr || index < 0 || index > this->size() - 1) {
            return;
        }
        Node<T> *temp = head;
        if (index == 0) {
            head = head->next;
            head->prev = nullptr;
            delete temp;
        } else {
            for (int i = 0; i < index - 1; i++) {
                temp = temp->next;
            }
            Node<T> *newTemp = temp->next;
            if(index == this->size() - 1) {
                temp->next = nullptr;
            } else {
                temp->next->next->prev = temp;
                temp->next = temp->next->next;
            }
            delete newTemp;
        }
    }
    void reverse() {
        Node<T> *temp = nullptr, *curr = head;
        while(curr != nullptr) {
            temp = curr->prev;
            curr->prev = curr->next;
            curr->next = temp;
            curr = curr->prev;
        }
        if (temp != nullptr) {
            head = temp->prev;
        }
    }
    void print() {
        Node<T> *temp = head;
        while(temp->next != nullptr) {
            std::cout << temp->data << " <=> ";
            temp = temp->next;
        }
        std::cout << temp->data << std::endl;
    }
};

int main() {
    DoublyLinkedList<int> dll;
}

